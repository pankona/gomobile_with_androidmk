all:
	cd $(CURDIR) && ndk-build
	make -C $(CURDIR)/go

clean:
	cd $(CURDIR) && ndk-build clean
	make clean -C $(CURDIR)/go
