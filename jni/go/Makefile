OUTDIR = $(CURDIR)/../../libs
AAR    = $(OUTDIR)/hello.aar

all:
	mkdir -p $(OUTDIR)
	CGO_LDFLAGS_ALLOW=".*" gomobile bind -target=android -v -o $(AAR)

clean:
	rm -rf $(OUTDIR)/*
