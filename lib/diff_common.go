package jd

func diff(
	a, b JsonNode,
	p path,
	metadata []Metadata,
	strategy patchStrategy,
) Diff {
	d := make(Diff, 0)
	if b.Equals(a, metadata...) {
		return d
	}
	var de DiffElement
	switch strategy {
	case mergePatchStrategy:
		de = DiffElement{
			Path:      p.prependMetadataMerge(),
			NewValues: jsonArray{a},
		}
	default:
		de = DiffElement{
			Path:      p.clone(),
			OldValues: nodeList(b),
			NewValues: nodeList(a),
		}
	}
	return append(d, de)
}
