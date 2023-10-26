package generator_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/mod/modfile"

	"github.com/mercari/grpc-federation/generator"
	"github.com/mercari/grpc-federation/internal/testutil"
	"github.com/mercari/grpc-federation/resolver"
)

func TestCodeGenerate(t *testing.T) {
	tmpDir := filepath.Join(t.TempDir(), "grpc-federation")

	tests := []string{
		"simple_aggregation",
		"minimum",
		"create_post",
		"custom_resolver",
		"async",
		"alias",
		"autobind",
		"const_value",
		"multi_user",
		"resolver_overlaps",
		"oneof",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			files := testutil.Compile(t, filepath.Join(testutil.RepoRoot(), "testdata", test+".proto"))

			var dependentFiles []string
			for _, file := range files {
				if strings.Contains(file.GetName(), "/") {
					continue
				}
				dependentFiles = append(dependentFiles, file.GetName())
			}

			r := resolver.New(files)
			result, err := r.Resolve()
			if err != nil {
				t.Fatal(err)
			}
			if len(result.Services) != 1 {
				t.Fatalf("faield to get services. expected 1 but got %d", len(result.Services))
			}
			service := result.Services[0]
			out, err := generator.NewCodeGenerator().Generate(service)
			if err != nil {
				t.Fatal(err)
			}
			path := filepath.Join("testdata", fmt.Sprintf("expected_%s.go", test))
			data, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(string(out), string(data)); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}

			// Tests whether the automatically generated files can be compiled.
			// Compilation also requires files generated by `protoc-gen-go` and `protoc-gen-go-grpc`, so we generate them and build them.
			t.Run("build test", func(t *testing.T) {
				var srcFiles []string
				for _, dependentFile := range dependentFiles {
					srcFiles = append(srcFiles, filepath.Join("..", "testdata", dependentFile))
				}
				content, err := yaml.Marshal(struct {
					Imports []string                  `yaml:"imports"`
					Src     []string                  `yaml:"src"`
					Out     string                    `yaml:"out"`
					Plugins []*generator.PluginConfig `yaml:"plugins"`
				}{
					Imports: []string{filepath.Join("..", "testdata")},
					Src:     srcFiles,
					Out:     tmpDir,
					Plugins: []*generator.PluginConfig{
						{Plugin: "go", Opt: "paths=import"},
						{Plugin: "go-grpc", Opt: "paths=import"},
						{Plugin: "grpc-federation", Opt: "paths=import"},
					},
				})
				if err != nil {
					t.Fatal(err)
				}
				cfg, err := generator.LoadConfigFromReader(bytes.NewBuffer(content))
				if err != nil {
					t.Fatal(err)
				}
				g := generator.New(cfg)
				protoToRespMap, err := g.GenerateAll(context.Background())
				if err != nil {
					t.Fatal(err)
				}

				// TODO: The current implementation expects the `go_package` defined in the proto file of testdata to start with "example/".
				// To support other packages, this process needs to be modified.
				modFile := new(modfile.File)
				if err := modFile.AddModuleStmt("example"); err != nil {
					t.Fatal(err)
				}
				if err := modFile.AddGoStmt("1.21"); err != nil {
					t.Fatal(err)
				}
				modContent, err := modFile.Format()
				if err != nil {
					t.Fatal(err)
				}
				modFilePath := filepath.Join(tmpDir, test, "example", "go.mod")
				t.Logf("write %s", modFilePath)
				if err := os.MkdirAll(filepath.Dir(modFilePath), 0o755); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(modFilePath, modContent, 0o600); err != nil {
					t.Fatal(err)
				}
				var federationFilePath string
				for _, responses := range protoToRespMap {
					for _, response := range responses {
						for _, file := range response.File {
							path := filepath.Join(tmpDir, test, file.GetName())
							if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
								t.Fatal(err)
							}
							t.Logf("write %s", path)
							if strings.HasSuffix(path, "_grpc_federation.go") {
								federationFilePath = path
								if err := os.WriteFile(path, out, 0o600); err != nil {
									t.Fatal(err)
								}
							} else {
								if err := os.WriteFile(path, []byte(file.GetContent()), 0o600); err != nil {
									t.Fatal(err)
								}
							}
						}
					}
				}
				if federationFilePath == "" {
					t.Fatalf("failed to find grpc federation file")
				}
				buildRelPath, err := filepath.Rel(filepath.Dir(modFilePath), filepath.Dir(federationFilePath))
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("build relative path: %s", buildRelPath)
				t.Run("go mod tidy", func(t *testing.T) {
					cmd := exec.Command("go", "mod", "tidy")
					cmd.Dir = filepath.Dir(modFilePath)
					out, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatalf("%q: %s", out, err)
					}
				})
				t.Run("go build", func(t *testing.T) {
					cmd := exec.Command("go", "build", fmt.Sprintf("./%s", buildRelPath)) //nolint: gosec
					cmd.Dir = filepath.Dir(modFilePath)
					out, err := cmd.CombinedOutput()
					if err != nil {
						t.Fatalf("%q: %s", out, err)
					}
				})
			})
		})
	}
}
