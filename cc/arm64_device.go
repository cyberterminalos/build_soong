package cc

import (
	"strings"

	"android/soong/common"
)

var (
	arm64Cflags = []string{
		"-fno-exceptions", // from build/core/combo/select.mk
		"-Wno-multichar",  // from build/core/combo/select.mk
		"-fno-strict-aliasing",
		"-fstack-protector",
		"-ffunction-sections",
		"-fdata-sections",
		"-funwind-tables",
		"-Wa,--noexecstack",
		"-Werror=format-security",
		"-D_FORTIFY_SOURCE=2",
		"-fno-short-enums",
		"-no-canonical-prefixes",
		"-fno-canonical-system-headers",
		"-include ${SrcDir}/build/core/combo/include/arch/linux-arm64/AndroidConfig.h",

		// Help catch common 32/64-bit errors.
		"-Werror=pointer-to-int-cast",
		"-Werror=int-to-pointer-cast",

		"-fno-strict-volatile-bitfields",

		// TARGET_RELEASE_CFLAGS
		"-DNDEBUG",
		"-O2 -g",
		"-Wstrict-aliasing=2",
		"-fgcse-after-reload",
		"-frerun-cse-after-loop",
		"-frename-registers",
	}

	arm64Ldflags = []string{
		"-Wl,-z,noexecstack",
		"-Wl,-z,relro",
		"-Wl,-z,now",
		"-Wl,--build-id=md5",
		"-Wl,--warn-shared-textrel",
		"-Wl,--fatal-warnings",
		"-Wl,-maarch64linux",
		"-Wl,--hash-style=gnu",

		// Disable transitive dependency library symbol resolving.
		"-Wl,--allow-shlib-undefined",
	}

	arm64Cppflags = []string{
		"-fvisibility-inlines-hidden",
	}
)

func init() {
	pctx.StaticVariable("arm64GccVersion", "4.9")

	pctx.StaticVariable("arm64GccRoot",
		"${SrcDir}/prebuilts/gcc/${HostPrebuiltTag}/aarch64/aarch64-linux-android-${arm64GccVersion}")

	pctx.StaticVariable("arm64GccTriple", "aarch64-linux-android")

	pctx.StaticVariable("arm64Cflags", strings.Join(arm64Cflags, " "))
	pctx.StaticVariable("arm64Ldflags", strings.Join(arm64Ldflags, " "))
	pctx.StaticVariable("arm64Cppflags", strings.Join(arm64Cppflags, " "))
	pctx.StaticVariable("arm64IncludeFlags", strings.Join([]string{
		"-isystem ${LibcRoot}/arch-arm64/include",
		"-isystem ${LibcRoot}/include",
		"-isystem ${LibcRoot}/kernel/uapi",
		"-isystem ${LibcRoot}/kernel/uapi/asm-arm64",
		"-isystem ${LibmRoot}/include",
		"-isystem ${LibmRoot}/include/arm64",
	}, " "))

	pctx.StaticVariable("arm64ClangCflags", strings.Join(clangFilterUnknownCflags(arm64Cflags), " "))
	pctx.StaticVariable("arm64ClangLdflags", strings.Join(clangFilterUnknownCflags(arm64Ldflags), " "))
	pctx.StaticVariable("arm64ClangCppflags", strings.Join(clangFilterUnknownCflags(arm64Cppflags), " "))
}

type toolchainArm64 struct {
	toolchain64Bit
}

var toolchainArm64Singleton = &toolchainArm64{}

func (t *toolchainArm64) Name() string {
	return "arm64"
}

func (t *toolchainArm64) GccRoot() string {
	return "${arm64GccRoot}"
}

func (t *toolchainArm64) GccTriple() string {
	return "${arm64GccTriple}"
}

func (t *toolchainArm64) GccVersion() string {
	return "${arm64GccVersion}"
}

func (t *toolchainArm64) Cflags() string {
	return "${arm64Cflags} ${arm64IncludeFlags}"
}

func (t *toolchainArm64) Cppflags() string {
	return "${arm64Cppflags}"
}

func (t *toolchainArm64) Ldflags() string {
	return "${arm64Ldflags}"
}

func (t *toolchainArm64) IncludeFlags() string {
	return "${arm64IncludeFlags}"
}

func (t *toolchainArm64) ClangTriple() string {
	return "${arm64GccTriple}"
}

func (t *toolchainArm64) ClangCflags() string {
	return "${arm64ClangCflags}"
}

func (t *toolchainArm64) ClangCppflags() string {
	return "${arm64ClangCppflags}"
}

func (t *toolchainArm64) ClangLdflags() string {
	return "${arm64Ldflags}"
}

func arm64ToolchainFactory(archVariant string, cpuVariant string) Toolchain {
	return toolchainArm64Singleton
}

func init() {
	registerToolchainFactory(common.Device, common.Arm64, arm64ToolchainFactory)
}
