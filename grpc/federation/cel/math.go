package cel

import (
	"context"
	"math"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
)

const MathPackageName = "math"

var _ cel.SingletonLibrary = new(MathLibrary)

type MathLibrary struct {
}

func NewMathLibrary() *MathLibrary {
	return &MathLibrary{}
}

func (lib *MathLibrary) LibraryName() string {
	return packageName(MathPackageName)
}

func createMathName(name string) string {
	return createName(MathPackageName, name)
}

func createMathID(name string) string {
	return createID(MathPackageName, name)
}

func (lib *MathLibrary) CompileOptions() []cel.EnvOption {
	opts := []cel.EnvOption{}

	for _, funcOpts := range [][]cel.EnvOption{
		// math package functions
		BindFunction(
			createMathName("abs"),
			OverloadFunc(createMathID("abs_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Abs(float64(args[0].(types.Double))))
				},
			),
			OverloadFunc(createMathID("abs_int_double"), []*cel.Type{cel.IntType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Abs(float64(args[0].(types.Int))))
				},
			),
		),
		BindFunction(
			createMathName("acos"),
			OverloadFunc(createMathID("acos_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Acos(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("acosh"),
			OverloadFunc(createMathID("acosh_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Acosh(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("asin"),
			OverloadFunc(createMathID("asin_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Asin(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("asinh"),
			OverloadFunc(createMathID("asinh_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Asinh(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("atan"),
			OverloadFunc(createMathID("atan_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Atan(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("atan2"),
			OverloadFunc(createMathID("atan2_double_double_double"), []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Atan2(float64(args[0].(types.Double)), float64(args[1].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("atanh"),
			OverloadFunc(createMathID("atanh_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Atanh(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("cbrt"),
			OverloadFunc(createMathID("cbrt_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Cbrt(float64(args[0].(types.Double))))
				},
			),
			OverloadFunc(createMathID("cbrt_int_double"), []*cel.Type{cel.IntType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Cbrt(float64(args[0].(types.Int))))
				},
			),
		),
		BindFunction(
			createMathName("ceil"),
			OverloadFunc(createMathID("ceil_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Ceil(float64(args[0].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("copysign"),
			OverloadFunc(createMathID("copysign_double_doubles_double"), []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Copysign(float64(args[0].(types.Double)), float64(args[1].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("sqrt"),
			OverloadFunc(createMathID("sqrt_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Sqrt(float64(args[0].(types.Double))))
				},
			),
			OverloadFunc(createMathID("sqrt_int_double"), []*cel.Type{cel.IntType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Sqrt(float64(args[0].(types.Int))))
				},
			),
		),
		BindFunction(
			createMathName("pow"),
			OverloadFunc(createMathID("pow_double_double_double"), []*cel.Type{cel.DoubleType, cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Pow(float64(args[0].(types.Double)), float64(args[1].(types.Double))))
				},
			),
		),
		BindFunction(
			createMathName("floor"),
			OverloadFunc(createMathID("floor_double_double"), []*cel.Type{cel.DoubleType}, cel.DoubleType,
				func(_ context.Context, args ...ref.Val) ref.Val {
					return types.Double(math.Floor(float64(args[0].(types.Double))))
				},
			),
		),
	} {
		opts = append(opts, funcOpts...)
	}

	return opts
}

func (lib *MathLibrary) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{}
}
