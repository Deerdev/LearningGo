package lib

// internal代码包中声明的公开程序实体仅能被该代码包的直接父包及其子包中的 代码引用。
// 当然，引用前需要先导入这个internal包。对于其他代码包，导入该internal包 都是非法的，无法通过编译
import (
    "os"
	in "Base/04internal_package/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
