#include <stdio.h>
#include <android/log.h>

void print() {
    printf("hello world!");
    __android_log_print(ANDROID_LOG_INFO, "hello", "%s(%d)", __FILE__, __LINE__);
}
