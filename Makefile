all:
	make -C $(CURDIR)/jni

clean:
	make clean -C $(CURDIR)/jni
	rm -rf $(CURDIR)/libs
	rm -rf $(CURDIR)/obj
