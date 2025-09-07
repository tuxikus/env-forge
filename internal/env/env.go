package env

type Link struct {
	Src string
	Dst string
}

type Dir struct {
	Path string
}

type Pkg struct {
	Name string
}

type Env struct {
	Links []Link
	Dirs  []Dir
	Pkgs  []Pkg
}
