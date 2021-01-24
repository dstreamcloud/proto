// Code generated by protoc-gen-clone. DO NOT EDIT.

package dag_pb

func (z *V1_Release) Clone() *V1_Release {
	if z == nil {
		return nil
	}
	zz := &V1_Release{}
	zz.Id = z.Id
	zz.OrganizationID = z.OrganizationID
	zz.OwnerID = z.OwnerID
	zz.Name = z.Name
	zz.Remarks = z.Remarks
	zz.EntryID = z.EntryID
	zz.ExitID = z.ExitID
	zz0 := make([]*V1_Release, len(z.Modules))
	for i := range z.Modules {
		zz0[i] = z.Modules[i].Clone()
	}
	zz.Modules = zz0
	zz1 := make([]*V1_Node, len(z.Nodes))
	for i := range z.Nodes {
		zz1[i] = z.Nodes[i].Clone()
	}
	zz.Nodes = zz1
	zz2 := make([]*V1_Edge, len(z.Edges))
	for i := range z.Edges {
		zz2[i] = z.Edges[i].Clone()
	}
	zz.Edges = zz2
	zz3 := make(map[string]string)
	for k, v := range z.Inputs {
		zz3[k] = v
	}
	zz.Inputs = zz3
	return zz
}
func (z *V1_Source) Clone() *V1_Source {
	if z == nil {
		return nil
	}
	zz := &V1_Source{}
	zz.Id = z.Id
	zz.UpdatedAt = z.UpdatedAt
	zz.ProjectID = z.ProjectID
	zz.Name = z.Name
	zz.Remarks = z.Remarks
	zz.EntryID = z.EntryID
	zz0 := make([]string, len(z.Modules))
	for i := range z.Modules {
		zz0[i] = z.Modules[i]
	}
	zz.Modules = zz0
	zz1 := make([]*V1_Edge, len(z.Edges))
	for i := range z.Edges {
		zz1[i] = z.Edges[i].Clone()
	}
	zz.Edges = zz1
	zz2 := make(map[string]string)
	for k, v := range z.Inputs {
		zz2[k] = v
	}
	zz.Inputs = zz2
	return zz
}
func (z *V1_Node) Clone() *V1_Node {
	if z == nil {
		return nil
	}
	zz := &V1_Node{}
	zz.Id = z.Id
	zz.Type = z.Type
	zz.Data = z.Data
	zz0 := make(map[string]string)
	for k, v := range z.Inputs {
		zz0[k] = v
	}
	zz.Inputs = zz0
	return zz
}
func (z *V1_Edge) Clone() *V1_Edge {
	if z == nil {
		return nil
	}
	zz := &V1_Edge{}
	zz.Id = z.Id
	zz.Source = z.Source
	zz.Target = z.Target
	return zz
}
func (z *V1) Clone() *V1 {
	if z == nil {
		return nil
	}
	zz := &V1{}
	return zz
}