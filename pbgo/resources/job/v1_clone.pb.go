// Code generated by protoc-gen-clone. DO NOT EDIT.

package job_pb

func (z *V1_Pipeline) Clone() *V1_Pipeline {
	if z == nil {
		return nil
	}
	zz := &V1_Pipeline{}
	zz.Id = z.Id
	zz.OrganizationID = z.OrganizationID
	zz.OwnerID = z.OwnerID
	zz.Settings = z.Settings.Clone()
	zz0 := make(map[string]string)
	for k, v := range z.Arguments {
		zz0[k] = v
	}
	zz.Arguments = zz0
	zz1 := make(map[string]string)
	for k, v := range z.Environments {
		zz1[k] = v
	}
	zz.Environments = zz1
	return zz
}
func (z *V1_Job) Clone() *V1_Job {
	if z == nil {
		return nil
	}
	zz := &V1_Job{}
	zz.Id = z.Id
	zz.OrganizationID = z.OrganizationID
	zz.OwnerID = z.OwnerID
	zz.AgentID = z.AgentID
	zz.Settings = z.Settings.Clone()
	zz0 := make(map[string]string)
	for k, v := range z.Arguments {
		zz0[k] = v
	}
	zz.Arguments = zz0
	zz1 := make(map[string]string)
	for k, v := range z.Environments {
		zz1[k] = v
	}
	zz.Environments = zz1
	return zz
}
func (z *V1_Settings) Clone() *V1_Settings {
	if z == nil {
		return nil
	}
	zz := &V1_Settings{}
	zz.Concurrency = z.Concurrency
	zz.Total = z.Total
	zz.StopAtUs = z.StopAtUs
	return zz
}
func (z *V1) Clone() *V1 {
	if z == nil {
		return nil
	}
	zz := &V1{}
	return zz
}
