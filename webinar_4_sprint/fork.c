#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

#define MAX_PROCESSES 100

int main(int argc, char** argv) {
    pid_t* pids = calloc(MAX_PROCESSES, sizeof(pid_t));
    int is_parent = 0;
    for (int i = 0; i < MAX_PROCESSES; ++i) {
        pid_t pid = fork();
        if (pid != 0) {
            pids[i] = pid;
        } else {
            is_parent = 1;
        }
    }
    if (is_parent) {
        for (int i = 0; i < 100; ++i) {
            waitpid(pids[i], 0, 0);
        }
    }
    return 0;
}
