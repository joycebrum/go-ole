// +build !windows

package ole

func (enum *IEnumVARIANT) Clone() (*IEnumVARIANT, error) {
	return nil, NewError(E_NOTIMPL)
}

func (enum *IEnumVARIANT) Reset() error {
	return NewError(E_NOTIMPL)
}

func (enum *IEnumVARIANT) Skip(celt uint) error {
	return NewError(E_NOTIMPL)
}

func (enum *IEnumVARIANT) Next(celt uint) (array VARIANT, length uint, err error) {
	return nil, 0, NewError(E_NOTIMPL)
}
