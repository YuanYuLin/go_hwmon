import ops
import ops_git
import ops_golang
import iopc

pkg_path=""
output_dir=""
GOPATH=""
go_pkgs_dep_script="go_dep_packages.sh"
go_pkgs_build_script="go_build_packages.sh"
servlet_dir="build"

def set_global(args):
    global pkg_path
    global output_dir
    global GOPATH
    global servlet_dir
    arch = ops.getEnv("ARCH_ALT")
    pkg_path = args["pkg_path"]
    output_dir = args["output_path"]
    GOPATH=output_dir
    servlet_dir = ops.path_join(output_dir, "/build")

def MAIN_ENV(args):
    set_global(args)

    print ops.getEnv("GOROOT")
    print ops.getEnv("PATH")
    ops.exportEnv(ops.setEnv("GOPATH", GOPATH))
    #ops_golang.get(GOPATH, "github.com/gorilla/mux")

    return False

def MAIN_EXTRACT(args):
    set_global(args)

    CMD = [ops.path_join(pkg_path, go_pkgs_dep_script)]
    res = ops.execCmd(CMD, GOPATH, False, None)
    #if res[2] != 0:
    #    print res
    #    print res[1]
    #    sys.exit(1)

    ops.copyto(ops.path_join(pkg_path, "src"), GOPATH)
    ops.copyto(ops.path_join(pkg_path, "Makefile"), GOPATH)
    ops.copyto(ops.path_join(pkg_path, "main.go"), GOPATH)

    return True

def MAIN_PATCH(args, patch_group_name):
    set_global(args)
    for patch in iopc.get_patch_list(pkg_path, patch_group_name):
        if iopc.apply_patch(tarball_dir, patch):
            continue
        else:
            sys.exit(1)

    return True

def MAIN_CONFIGURE(args):
    set_global(args)

    return True

def MAIN_BUILD(args):
    set_global(args)

    '''
    CMD = [ops.path_join(pkg_path, go_pkgs_build_script)]
    res = ops.execCmd(CMD, GOPATH, False, None)
    if res[2] != 0:
        print res
        print res[1]
        sys.exit(1)
    '''
    iopc.make(GOPATH)
    #iopc.make(src_path)

    return False

def MAIN_INSTALL(args):
    set_global(args)

    iopc.installBin(args["pkg_name"], ops.path_join(servlet_dir, "."), "www")

    return False

def MAIN_SDKENV(args):
    set_global(args)

    return False

def MAIN_CLEAN_BUILD(args):
    set_global(args)

    return False

def MAIN(args):
    set_global(args)

