package goroutineAffinity

import (
	"fmt"
	"runtime"
)

/*
#if __APPLE__
    int ga_set_affinity(int cpuid) {
        return -1;
    }
    int ga_get_affinity_cpu() {
        return -1;
    }
#elif defined(WIN32) || defined(_WIN32) || defined(__WIN32__) || defined(__NT__) || defined(_WIN64) || defined(__WINDOWS__)
    int ga_set_affinity(int cpuid) {
        return -1;
    }
    int ga_get_affinity_cpu() {
        return -1;
    }
#elif defined(__LINUX__) || defined(linux) || defined(__linux)
    #define _GNU_SOURCE
    #include <sched.h>
    #include <pthread.h>
    int ga_set_affinity(int cpuid) {
        pthread_t thread_id;
        cpu_set_t cpuset;
        thread_id = pthread_self();
        CPU_ZERO(&cpuset);
        CPU_SET(cpuid, &cpuset);
        return pthread_setaffinity_np(thread_id, sizeof(cpu_set_t), &cpuset);
    }
    int ga_get_affinity_cpu() {
        return sched_getcpu();
    }
#elif defined(__FreeBSD__) || defined(__FreeBSD_kernel__)
    int ga_set_affinity(int cpuid) {
        return -1;
    }
    int ga_get_affinity_cpu() {
        return -1;
    }
#else
    int ga_set_affinity(int cpuid) {
        return -1;
    }
    int ga_get_affinity_cpu() {
        return -1;
    }
#endif
*/
import "C"

func isSupport() (string, bool) {
	os := runtime.GOOS
	switch os {
	case "windows":
		return "windows", false
	case "darwin":
		return "darwin/OSX", false
	case "linux":
		return "linux", true
	default:
		return "unknown", false
	}
}

func SetAffinity(cpuID int) int {
	if osName, ok := isSupport(); !ok {
		fmt.Println("goroutineAffinity say OS:", osName, "is unsupported")
		return -1
	}

	runtime.LockOSThread()
	return int(C.ga_set_affinity(C.int(cpuID)))
}

func GetAffinityCPU() int {
	if osName, ok := isSupport(); !ok {
		fmt.Println("goroutineAffinity say OS:", osName, "is unsupported")
		return -1
	}

	return int(C.ga_get_affinity_cpu())
}
