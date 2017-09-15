import re

def gen():
	protolist = []
	rec = re.compile("message\s(\w+)\s{")
	with open("./pb.proto") as f:
		for cell in f.readlines():
			s = re.match(rec, cell)
			if s:
				strs = s.group(1).strip()
				if strs.endswith("S2C") or strs.endswith("C2S"):
					formatstr = "p.proto2fun[settings.%s] = handler.%s" % (strs, strs)
					print formatstr
					protolist.append(s.group(1).strip())

	return protolist

l = gen()
# print l
