// generated file; DO NOT EDIT - use go generate in directory with source

// +build 386,openbsd

package syscallpp

func GetName(num int) string {
	switch num {
	case 0:
		return "exit"
	case 1:
		return "fork"
	case 2:
		return "read"
	case 3:
		return "write"
	case 4:
		return "open"
	case 5:
		return "close"
	case 6:
		return "__tfork"
	case 7:
		return "link"
	case 8:
		return "unlink"
	case 9:
		return "wait4"
	case 10:
		return "chdir"
	case 11:
		return "fchdir"
	case 12:
		return "mknod"
	case 13:
		return "chmod"
	case 14:
		return "chown"
	case 15:
		return "obreak"
	case 16:
		return "getdtablecount"
	case 17:
		return "getrusage"
	case 18:
		return "getpid"
	case 19:
		return "mount"
	case 20:
		return "unmount"
	case 21:
		return "setuid"
	case 22:
		return "getuid"
	case 23:
		return "geteuid"
	case 24:
		return "ptrace"
	case 25:
		return "recvmsg"
	case 26:
		return "sendmsg"
	case 27:
		return "recvfrom"
	case 28:
		return "accept"
	case 29:
		return "getpeername"
	case 30:
		return "getsockname"
	case 31:
		return "access"
	case 32:
		return "chflags"
	case 33:
		return "fchflags"
	case 34:
		return "sync"
	case 35:
		return "kill"
	case 36:
		return "stat"
	case 37:
		return "getppid"
	case 38:
		return "lstat"
	case 39:
		return "dup"
	case 40:
		return "fstatat"
	case 41:
		return "getegid"
	case 42:
		return "profil"
	case 43:
		return "ktrace"
	case 44:
		return "sigaction"
	case 45:
		return "getgid"
	case 46:
		return "sigprocmask"
	case 47:
		return "getlogin"
	case 48:
		return "setlogin"
	case 49:
		return "acct"
	case 50:
		return "sigpending"
	case 51:
		return "fstat"
	case 52:
		return "ioctl"
	case 53:
		return "reboot"
	case 54:
		return "revoke"
	case 55:
		return "symlink"
	case 56:
		return "readlink"
	case 57:
		return "execve"
	case 58:
		return "umask"
	case 59:
		return "chroot"
	case 60:
		return "getfsstat"
	case 61:
		return "statfs"
	case 62:
		return "fstatfs"
	case 63:
		return "fhstatfs"
	case 64:
		return "vfork"
	case 65:
		return "gettimeofday"
	case 66:
		return "settimeofday"
	case 67:
		return "setitimer"
	case 68:
		return "getitimer"
	case 69:
		return "select"
	case 70:
		return "kevent"
	case 71:
		return "munmap"
	case 72:
		return "mprotect"
	case 73:
		return "madvise"
	case 74:
		return "utimes"
	case 75:
		return "futimes"
	case 76:
		return "mincore"
	case 77:
		return "getgroups"
	case 78:
		return "setgroups"
	case 79:
		return "getpgrp"
	case 80:
		return "setpgid"
	case 81:
		return "utimensat"
	case 82:
		return "futimens"
	case 83:
		return "clock_gettime"
	case 84:
		return "clock_settime"
	case 85:
		return "clock_getres"
	case 86:
		return "dup2"
	case 87:
		return "nanosleep"
	case 88:
		return "fcntl"
	case 89:
		return "__thrsleep"
	case 90:
		return "fsync"
	case 91:
		return "setpriority"
	case 92:
		return "socket"
	case 93:
		return "connect"
	case 94:
		return "getdents"
	case 95:
		return "getpriority"
	case 96:
		return "sigreturn"
	case 97:
		return "bind"
	case 98:
		return "setsockopt"
	case 99:
		return "listen"
	case 100:
		return "ppoll"
	case 101:
		return "pselect"
	case 102:
		return "sigsuspend"
	case 103:
		return "getsockopt"
	case 104:
		return "readv"
	case 105:
		return "writev"
	case 106:
		return "fchown"
	case 107:
		return "fchmod"
	case 108:
		return "setreuid"
	case 109:
		return "setregid"
	case 110:
		return "rename"
	case 111:
		return "flock"
	case 112:
		return "mkfifo"
	case 113:
		return "sendto"
	case 114:
		return "shutdown"
	case 115:
		return "socketpair"
	case 116:
		return "mkdir"
	case 117:
		return "rmdir"
	case 118:
		return "adjtime"
	case 119:
		return "setsid"
	case 120:
		return "quotactl"
	case 121:
		return "nfssvc"
	case 122:
		return "getfh"
	case 123:
		return "sysarch"
	case 124:
		return "pread"
	case 125:
		return "pwrite"
	case 126:
		return "setgid"
	case 127:
		return "setegid"
	case 128:
		return "seteuid"
	case 129:
		return "pathconf"
	case 130:
		return "fpathconf"
	case 131:
		return "swapctl"
	case 132:
		return "getrlimit"
	case 133:
		return "setrlimit"
	case 134:
		return "mmap"
	case 135:
		return "lseek"
	case 136:
		return "truncate"
	case 137:
		return "ftruncate"
	case 138:
		return "__sysctl"
	case 139:
		return "mlock"
	case 140:
		return "munlock"
	case 141:
		return "getpgid"
	case 142:
		return "utrace"
	case 143:
		return "semget"
	case 144:
		return "msgget"
	case 145:
		return "msgsnd"
	case 146:
		return "msgrcv"
	case 147:
		return "shmat"
	case 148:
		return "shmdt"
	case 149:
		return "minherit"
	case 150:
		return "poll"
	case 151:
		return "issetugid"
	case 152:
		return "lchown"
	case 153:
		return "getsid"
	case 154:
		return "msync"
	case 155:
		return "pipe"
	case 156:
		return "fhopen"
	case 157:
		return "preadv"
	case 158:
		return "pwritev"
	case 159:
		return "kqueue"
	case 160:
		return "mlockall"
	case 161:
		return "munlockall"
	case 162:
		return "getresuid"
	case 163:
		return "setresuid"
	case 164:
		return "getresgid"
	case 165:
		return "setresgid"
	case 166:
		return "mquery"
	case 167:
		return "closefrom"
	case 168:
		return "sigaltstack"
	case 169:
		return "shmget"
	case 170:
		return "semop"
	case 171:
		return "fhstat"
	case 172:
		return "__semctl"
	case 173:
		return "shmctl"
	case 174:
		return "msgctl"
	case 175:
		return "sched_yield"
	case 176:
		return "getthrid"
	case 177:
		return "__thrwakeup"
	case 178:
		return "__threxit"
	case 179:
		return "__thrsigdivert"
	case 180:
		return "__getcwd"
	case 181:
		return "adjfreq"
	case 182:
		return "setrtable"
	case 183:
		return "getrtable"
	case 184:
		return "faccessat"
	case 185:
		return "fchmodat"
	case 186:
		return "fchownat"
	case 187:
		return "linkat"
	case 188:
		return "mkdirat"
	case 189:
		return "mkfifoat"
	case 190:
		return "mknodat"
	case 191:
		return "openat"
	case 192:
		return "readlinkat"
	case 193:
		return "renameat"
	case 194:
		return "symlinkat"
	case 195:
		return "unlinkat"
	case 196:
		return "__set_tcb"
	case 197:
		return "__get_tcb"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "exit":
		return 0
	case "fork":
		return 1
	case "read":
		return 2
	case "write":
		return 3
	case "open":
		return 4
	case "close":
		return 5
	case "__tfork":
		return 6
	case "link":
		return 7
	case "unlink":
		return 8
	case "wait4":
		return 9
	case "chdir":
		return 10
	case "fchdir":
		return 11
	case "mknod":
		return 12
	case "chmod":
		return 13
	case "chown":
		return 14
	case "obreak":
		return 15
	case "getdtablecount":
		return 16
	case "getrusage":
		return 17
	case "getpid":
		return 18
	case "mount":
		return 19
	case "unmount":
		return 20
	case "setuid":
		return 21
	case "getuid":
		return 22
	case "geteuid":
		return 23
	case "ptrace":
		return 24
	case "recvmsg":
		return 25
	case "sendmsg":
		return 26
	case "recvfrom":
		return 27
	case "accept":
		return 28
	case "getpeername":
		return 29
	case "getsockname":
		return 30
	case "access":
		return 31
	case "chflags":
		return 32
	case "fchflags":
		return 33
	case "sync":
		return 34
	case "kill":
		return 35
	case "stat":
		return 36
	case "getppid":
		return 37
	case "lstat":
		return 38
	case "dup":
		return 39
	case "fstatat":
		return 40
	case "getegid":
		return 41
	case "profil":
		return 42
	case "ktrace":
		return 43
	case "sigaction":
		return 44
	case "getgid":
		return 45
	case "sigprocmask":
		return 46
	case "getlogin":
		return 47
	case "setlogin":
		return 48
	case "acct":
		return 49
	case "sigpending":
		return 50
	case "fstat":
		return 51
	case "ioctl":
		return 52
	case "reboot":
		return 53
	case "revoke":
		return 54
	case "symlink":
		return 55
	case "readlink":
		return 56
	case "execve":
		return 57
	case "umask":
		return 58
	case "chroot":
		return 59
	case "getfsstat":
		return 60
	case "statfs":
		return 61
	case "fstatfs":
		return 62
	case "fhstatfs":
		return 63
	case "vfork":
		return 64
	case "gettimeofday":
		return 65
	case "settimeofday":
		return 66
	case "setitimer":
		return 67
	case "getitimer":
		return 68
	case "select":
		return 69
	case "kevent":
		return 70
	case "munmap":
		return 71
	case "mprotect":
		return 72
	case "madvise":
		return 73
	case "utimes":
		return 74
	case "futimes":
		return 75
	case "mincore":
		return 76
	case "getgroups":
		return 77
	case "setgroups":
		return 78
	case "getpgrp":
		return 79
	case "setpgid":
		return 80
	case "utimensat":
		return 81
	case "futimens":
		return 82
	case "clock_gettime":
		return 83
	case "clock_settime":
		return 84
	case "clock_getres":
		return 85
	case "dup2":
		return 86
	case "nanosleep":
		return 87
	case "fcntl":
		return 88
	case "__thrsleep":
		return 89
	case "fsync":
		return 90
	case "setpriority":
		return 91
	case "socket":
		return 92
	case "connect":
		return 93
	case "getdents":
		return 94
	case "getpriority":
		return 95
	case "sigreturn":
		return 96
	case "bind":
		return 97
	case "setsockopt":
		return 98
	case "listen":
		return 99
	case "ppoll":
		return 100
	case "pselect":
		return 101
	case "sigsuspend":
		return 102
	case "getsockopt":
		return 103
	case "readv":
		return 104
	case "writev":
		return 105
	case "fchown":
		return 106
	case "fchmod":
		return 107
	case "setreuid":
		return 108
	case "setregid":
		return 109
	case "rename":
		return 110
	case "flock":
		return 111
	case "mkfifo":
		return 112
	case "sendto":
		return 113
	case "shutdown":
		return 114
	case "socketpair":
		return 115
	case "mkdir":
		return 116
	case "rmdir":
		return 117
	case "adjtime":
		return 118
	case "setsid":
		return 119
	case "quotactl":
		return 120
	case "nfssvc":
		return 121
	case "getfh":
		return 122
	case "sysarch":
		return 123
	case "pread":
		return 124
	case "pwrite":
		return 125
	case "setgid":
		return 126
	case "setegid":
		return 127
	case "seteuid":
		return 128
	case "pathconf":
		return 129
	case "fpathconf":
		return 130
	case "swapctl":
		return 131
	case "getrlimit":
		return 132
	case "setrlimit":
		return 133
	case "mmap":
		return 134
	case "lseek":
		return 135
	case "truncate":
		return 136
	case "ftruncate":
		return 137
	case "__sysctl":
		return 138
	case "mlock":
		return 139
	case "munlock":
		return 140
	case "getpgid":
		return 141
	case "utrace":
		return 142
	case "semget":
		return 143
	case "msgget":
		return 144
	case "msgsnd":
		return 145
	case "msgrcv":
		return 146
	case "shmat":
		return 147
	case "shmdt":
		return 148
	case "minherit":
		return 149
	case "poll":
		return 150
	case "issetugid":
		return 151
	case "lchown":
		return 152
	case "getsid":
		return 153
	case "msync":
		return 154
	case "pipe":
		return 155
	case "fhopen":
		return 156
	case "preadv":
		return 157
	case "pwritev":
		return 158
	case "kqueue":
		return 159
	case "mlockall":
		return 160
	case "munlockall":
		return 161
	case "getresuid":
		return 162
	case "setresuid":
		return 163
	case "getresgid":
		return 164
	case "setresgid":
		return 165
	case "mquery":
		return 166
	case "closefrom":
		return 167
	case "sigaltstack":
		return 168
	case "shmget":
		return 169
	case "semop":
		return 170
	case "fhstat":
		return 171
	case "__semctl":
		return 172
	case "shmctl":
		return 173
	case "msgctl":
		return 174
	case "sched_yield":
		return 175
	case "getthrid":
		return 176
	case "__thrwakeup":
		return 177
	case "__threxit":
		return 178
	case "__thrsigdivert":
		return 179
	case "__getcwd":
		return 180
	case "adjfreq":
		return 181
	case "setrtable":
		return 182
	case "getrtable":
		return 183
	case "faccessat":
		return 184
	case "fchmodat":
		return 185
	case "fchownat":
		return 186
	case "linkat":
		return 187
	case "mkdirat":
		return 188
	case "mkfifoat":
		return 189
	case "mknodat":
		return 190
	case "openat":
		return 191
	case "readlinkat":
		return 192
	case "renameat":
		return 193
	case "symlinkat":
		return 194
	case "unlinkat":
		return 195
	case "__set_tcb":
		return 196
	case "__get_tcb":
		return 197
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "exit":
		return []ArgType{ARG_INT}
	case "fork":
		return []ArgType(nil)
	case "read":
		return []ArgType{ARG_INT, ARG_PTR}
	case "write":
		return []ArgType{ARG_INT, ARG_PTR}
	case "open":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "close":
		return []ArgType{ARG_INT}
	case "__tfork":
		return []ArgType(nil)
	case "link":
		return []ArgType{ARG_STR, ARG_STR}
	case "unlink":
		return []ArgType{ARG_STR}
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "chdir":
		return []ArgType{ARG_STR}
	case "fchdir":
		return []ArgType{ARG_INT}
	case "mknod":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "chmod":
		return []ArgType{ARG_STR, ARG_INT}
	case "chown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "obreak":
		return []ArgType(nil)
	case "getdtablecount":
		return []ArgType(nil)
	case "getrusage":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpid":
		return []ArgType(nil)
	case "mount":
		return []ArgType(nil)
	case "unmount":
		return []ArgType{ARG_STR, ARG_INT}
	case "setuid":
		return []ArgType{ARG_INT}
	case "getuid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "ptrace":
		return []ArgType(nil)
	case "recvmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "sendmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "recvfrom":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "accept":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpeername":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getsockname":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "access":
		return []ArgType{ARG_STR, ARG_INT}
	case "chflags":
		return []ArgType{ARG_STR, ARG_INT}
	case "fchflags":
		return []ArgType{ARG_INT, ARG_INT}
	case "sync":
		return []ArgType(nil)
	case "kill":
		return []ArgType{ARG_INT, ARG_PTR}
	case "stat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "getppid":
		return []ArgType(nil)
	case "lstat":
		return []ArgType{ARG_STR, ARG_PTR}
	case "dup":
		return []ArgType{ARG_INT}
	case "fstatat":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "profil":
		return []ArgType(nil)
	case "ktrace":
		return []ArgType(nil)
	case "sigaction":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "sigprocmask":
		return []ArgType(nil)
	case "getlogin":
		return []ArgType(nil)
	case "setlogin":
		return []ArgType{ARG_STR}
	case "acct":
		return []ArgType(nil)
	case "sigpending":
		return []ArgType(nil)
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "ioctl":
		return []ArgType(nil)
	case "reboot":
		return []ArgType(nil)
	case "revoke":
		return []ArgType{ARG_STR}
	case "symlink":
		return []ArgType{ARG_STR, ARG_STR}
	case "readlink":
		return []ArgType{ARG_STR, ARG_PTR}
	case "execve":
		return []ArgType(nil)
	case "umask":
		return []ArgType{ARG_INT}
	case "chroot":
		return []ArgType{ARG_STR}
	case "getfsstat":
		return []ArgType{ARG_PTR, ARG_INT}
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "fhstatfs":
		return []ArgType(nil)
	case "vfork":
		return []ArgType(nil)
	case "gettimeofday":
		return []ArgType{ARG_PTR}
	case "settimeofday":
		return []ArgType{ARG_PTR}
	case "setitimer":
		return []ArgType(nil)
	case "getitimer":
		return []ArgType(nil)
	case "select":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_PTR, ARG_PTR}
	case "kevent":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "mprotect":
		return []ArgType(nil)
	case "madvise":
		return []ArgType(nil)
	case "utimes":
		return []ArgType{ARG_STR, ARG_PTR}
	case "futimes":
		return []ArgType{ARG_INT, ARG_PTR}
	case "mincore":
		return []ArgType(nil)
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpgrp":
		return []ArgType(nil)
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "utimensat":
		return []ArgType(nil)
	case "futimens":
		return []ArgType(nil)
	case "clock_gettime":
		return []ArgType(nil)
	case "clock_settime":
		return []ArgType(nil)
	case "clock_getres":
		return []ArgType(nil)
	case "dup2":
		return []ArgType{ARG_INT, ARG_INT}
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "fcntl":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "__thrsleep":
		return []ArgType(nil)
	case "fsync":
		return []ArgType{ARG_INT}
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "socket":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "connect":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getdents":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "sigreturn":
		return []ArgType(nil)
	case "bind":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "setsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "listen":
		return []ArgType{ARG_INT, ARG_INT}
	case "ppoll":
		return []ArgType(nil)
	case "pselect":
		return []ArgType(nil)
	case "sigsuspend":
		return []ArgType(nil)
	case "getsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "readv":
		return []ArgType(nil)
	case "writev":
		return []ArgType(nil)
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "setreuid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setregid":
		return []ArgType{ARG_INT, ARG_INT}
	case "rename":
		return []ArgType{ARG_STR, ARG_STR}
	case "flock":
		return []ArgType{ARG_INT, ARG_INT}
	case "mkfifo":
		return []ArgType{ARG_STR, ARG_INT}
	case "sendto":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "shutdown":
		return []ArgType{ARG_INT, ARG_INT}
	case "socketpair":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "mkdir":
		return []ArgType{ARG_STR, ARG_INT}
	case "rmdir":
		return []ArgType{ARG_STR}
	case "adjtime":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "setsid":
		return []ArgType(nil)
	case "quotactl":
		return []ArgType(nil)
	case "nfssvc":
		return []ArgType(nil)
	case "getfh":
		return []ArgType(nil)
	case "sysarch":
		return []ArgType(nil)
	case "pread":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "pwrite":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "setgid":
		return []ArgType{ARG_INT}
	case "setegid":
		return []ArgType{ARG_INT}
	case "seteuid":
		return []ArgType{ARG_INT}
	case "pathconf":
		return []ArgType{ARG_STR, ARG_INT}
	case "fpathconf":
		return []ArgType{ARG_INT, ARG_INT}
	case "swapctl":
		return []ArgType(nil)
	case "getrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "lseek":
		return []ArgType(nil)
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "__sysctl":
		return []ArgType(nil)
	case "mlock":
		return []ArgType(nil)
	case "munlock":
		return []ArgType(nil)
	case "getpgid":
		return []ArgType{ARG_INT}
	case "utrace":
		return []ArgType(nil)
	case "semget":
		return []ArgType(nil)
	case "msgget":
		return []ArgType(nil)
	case "msgsnd":
		return []ArgType(nil)
	case "msgrcv":
		return []ArgType(nil)
	case "shmat":
		return []ArgType(nil)
	case "shmdt":
		return []ArgType(nil)
	case "minherit":
		return []ArgType(nil)
	case "poll":
		return []ArgType(nil)
	case "issetugid":
		return []ArgType(nil)
	case "lchown":
		return []ArgType{ARG_STR, ARG_INT, ARG_INT}
	case "getsid":
		return []ArgType{ARG_INT}
	case "msync":
		return []ArgType(nil)
	case "pipe":
		return []ArgType{ARG_PTR}
	case "fhopen":
		return []ArgType(nil)
	case "preadv":
		return []ArgType(nil)
	case "pwritev":
		return []ArgType(nil)
	case "kqueue":
		return []ArgType(nil)
	case "mlockall":
		return []ArgType(nil)
	case "munlockall":
		return []ArgType(nil)
	case "getresuid":
		return []ArgType(nil)
	case "setresuid":
		return []ArgType(nil)
	case "getresgid":
		return []ArgType(nil)
	case "setresgid":
		return []ArgType(nil)
	case "mquery":
		return []ArgType(nil)
	case "closefrom":
		return []ArgType(nil)
	case "sigaltstack":
		return []ArgType(nil)
	case "shmget":
		return []ArgType(nil)
	case "semop":
		return []ArgType(nil)
	case "fhstat":
		return []ArgType(nil)
	case "__semctl":
		return []ArgType(nil)
	case "shmctl":
		return []ArgType(nil)
	case "msgctl":
		return []ArgType(nil)
	case "sched_yield":
		return []ArgType(nil)
	case "getthrid":
		return []ArgType(nil)
	case "__thrwakeup":
		return []ArgType(nil)
	case "__threxit":
		return []ArgType(nil)
	case "__thrsigdivert":
		return []ArgType(nil)
	case "__getcwd":
		return []ArgType(nil)
	case "adjfreq":
		return []ArgType(nil)
	case "setrtable":
		return []ArgType(nil)
	case "getrtable":
		return []ArgType(nil)
	case "faccessat":
		return []ArgType(nil)
	case "fchmodat":
		return []ArgType(nil)
	case "fchownat":
		return []ArgType(nil)
	case "linkat":
		return []ArgType(nil)
	case "mkdirat":
		return []ArgType(nil)
	case "mkfifoat":
		return []ArgType(nil)
	case "mknodat":
		return []ArgType(nil)
	case "openat":
		return []ArgType(nil)
	case "readlinkat":
		return []ArgType(nil)
	case "renameat":
		return []ArgType(nil)
	case "symlinkat":
		return []ArgType(nil)
	case "unlinkat":
		return []ArgType(nil)
	case "__set_tcb":
		return []ArgType(nil)
	case "__get_tcb":
		return []ArgType(nil)
	}
	return nil
}