import re,sys,os

def get_imports(envoy_path, f):
    try:
        with open(os.path.join(envoy_path,"api/",f)) as data:
            alldata = data.read()
            allimports = re.findall("import \"(.*)\";", alldata)
            imports = set(list(allimports))
            for i in allimports:
                imports = imports.union(get_imports(envoy_path, i))
            imports.add(f)
            return imports
    except FileNotFoundError as e:
        print("did not find path", f)
        return set()


def get_sorted_imports(envoy_path, f):
    return sorted(get_imports(envoy_path, f))

def import_and_copy(f):
    envoy_path = os.getenv("ENVOYPATH")
    if not envoy_path:
        raise Exception("please set ENVOYPATH to envoy's root folder")
    for i in get_sorted_imports(envoy_path, f):
        print("importing", i)
        path = os.path.join(envoy_path,"api/",i)
        if os.path.exists(path):
            folder = os.path.dirname(i)
            option = f'option go_package = "github.com/solo-io/gloo/projects/gloo/pkg/api/external/{folder}";'
            if os.getenv("COPY"):
                f = os.system
            else:
                print("Run these shell:")
                f = lambda x: print(x)
            gopath = os.getenv("GOPATH")
            if not gopath:
                gopath = os.getenv("HOME")+"/go"

            basedir = f"{gopath}/src/github.com/solo-io/gloo/projects/gloo/api/external/"
            f(f'mkdir -p {basedir}{folder}')
            dest = basedir +i
            f(f'cp {path} {dest}')
            f(f"echo '{option}' >> {dest}")
            f("echo 'import \"extproto/ext.proto\";' >> " + dest)
            f(f"echo 'option (extproto.hash_all) = true;' >> {dest}")
            f(f"echo 'option (extproto.clone_all) = true;' >> {dest}")
            f(f"echo 'option (extproto.equal_all) = true;' >> {dest}")

def main():
    if len(sys.argv) != 2:
        print("please run like so:")
        print(
            f"    [COPY=1] [GOPATH=...] ENVOYPATH=... {sys.argv[0]} path-to-proto-in-envoy"
        )
        print("for example, this will copy route_components.proto and its dependencies to $GOPATH/src/github.com/solo-io/gloo/projects/gloo/api/external from ~/sources/envoy/api:")
        print(
            f"    COPY=1 GOPATH=~/go ENVOYPATH=~/sources/envoy {sys.argv[0]} envoy/config/route/v3/route_components.proto"
        )
        os.abort()
    import_and_copy(sys.argv[1])


if __name__ == "__main__":
    main()