// Code generated by protoc-gen-clone. DO NOT EDIT.

package agent_pb

func (z *V1_Agent) Clone() *V1_Agent {
	if z == nil {
		return nil
	}
	zz := &V1_Agent{}
	zz.Id = z.Id
	zz.OrganizationID = z.OrganizationID
	zz.Name = z.Name
	zz.Token = z.Token
	zz.Status = z.Status
	return zz
}
func (z *V1) Clone() *V1 {
	if z == nil {
		return nil
	}
	zz := &V1{}
	return zz
}