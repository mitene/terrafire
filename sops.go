package terrafire

type SopsClient interface {
	DecryptFile(file string) (string, error)
}