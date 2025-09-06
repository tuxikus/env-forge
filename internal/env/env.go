package env

type Link struct {
	Src string
	Dst string
}

type Dir struct {
	path string
}

type Pkg struct {
	name string
}

type Env struct {
	Links []Link
	Dirs  []Dir
	Pkgs  []Pkg
}
