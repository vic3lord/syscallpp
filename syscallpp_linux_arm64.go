// generated file; DO NOT EDIT - use go generate in directory with source

// +build arm64,linux

package syscallpp

func GetName(num int) string {
	switch num {
	case 0:
		return "io_setup"
	case 1:
		return "io_destroy"
	case 2:
		return "io_submit"
	case 3:
		return "io_cancel"
	case 4:
		return "io_getevents"
	case 5:
		return "setxattr"
	case 6:
		return "lsetxattr"
	case 7:
		return "fsetxattr"
	case 8:
		return "getxattr"
	case 9:
		return "lgetxattr"
	case 10:
		return "fgetxattr"
	case 11:
		return "listxattr"
	case 12:
		return "llistxattr"
	case 13:
		return "flistxattr"
	case 14:
		return "removexattr"
	case 15:
		return "lremovexattr"
	case 16:
		return "fremovexattr"
	case 17:
		return "getcwd"
	case 18:
		return "lookup_dcookie"
	case 19:
		return "eventfd2"
	case 20:
		return "epoll_create1"
	case 21:
		return "epoll_ctl"
	case 22:
		return "epoll_pwait"
	case 23:
		return "dup"
	case 24:
		return "dup3"
	case 25:
		return "fcntl"
	case 26:
		return "inotify_init1"
	case 27:
		return "inotify_add_watch"
	case 28:
		return "inotify_rm_watch"
	case 29:
		return "ioctl"
	case 30:
		return "ioprio_set"
	case 31:
		return "ioprio_get"
	case 32:
		return "flock"
	case 33:
		return "mknodat"
	case 34:
		return "mkdirat"
	case 35:
		return "unlinkat"
	case 36:
		return "symlinkat"
	case 37:
		return "linkat"
	case 38:
		return "renameat"
	case 39:
		return "umount2"
	case 40:
		return "mount"
	case 41:
		return "pivot_root"
	case 42:
		return "nfsservctl"
	case 43:
		return "statfs"
	case 44:
		return "fstatfs"
	case 45:
		return "truncate"
	case 46:
		return "ftruncate"
	case 47:
		return "fallocate"
	case 48:
		return "faccessat"
	case 49:
		return "chdir"
	case 50:
		return "fchdir"
	case 51:
		return "chroot"
	case 52:
		return "fchmod"
	case 53:
		return "fchmodat"
	case 54:
		return "fchownat"
	case 55:
		return "fchown"
	case 56:
		return "openat"
	case 57:
		return "close"
	case 58:
		return "vhangup"
	case 59:
		return "pipe2"
	case 60:
		return "quotactl"
	case 61:
		return "getdents64"
	case 62:
		return "lseek"
	case 63:
		return "read"
	case 64:
		return "write"
	case 65:
		return "readv"
	case 66:
		return "writev"
	case 67:
		return "pread64"
	case 68:
		return "pwrite64"
	case 69:
		return "preadv"
	case 70:
		return "pwritev"
	case 71:
		return "sendfile"
	case 72:
		return "pselect6"
	case 73:
		return "ppoll"
	case 74:
		return "signalfd4"
	case 75:
		return "vmsplice"
	case 76:
		return "splice"
	case 77:
		return "tee"
	case 78:
		return "readlinkat"
	case 79:
		return "fstatat"
	case 80:
		return "fstat"
	case 81:
		return "sync"
	case 82:
		return "fsync"
	case 83:
		return "fdatasync"
	case 84:
		return "sync_file_range2"
	case 85:
		return "sync_file_range"
	case 86:
		return "timerfd_create"
	case 87:
		return "timerfd_settime"
	case 88:
		return "timerfd_gettime"
	case 89:
		return "utimensat"
	case 90:
		return "acct"
	case 91:
		return "capget"
	case 92:
		return "capset"
	case 93:
		return "personality"
	case 94:
		return "exit"
	case 95:
		return "exit_group"
	case 96:
		return "waitid"
	case 97:
		return "set_tid_address"
	case 98:
		return "unshare"
	case 99:
		return "futex"
	case 100:
		return "set_robust_list"
	case 101:
		return "get_robust_list"
	case 102:
		return "nanosleep"
	case 103:
		return "getitimer"
	case 104:
		return "setitimer"
	case 105:
		return "kexec_load"
	case 106:
		return "init_module"
	case 107:
		return "delete_module"
	case 108:
		return "timer_create"
	case 109:
		return "timer_gettime"
	case 110:
		return "timer_getoverrun"
	case 111:
		return "timer_settime"
	case 112:
		return "timer_delete"
	case 113:
		return "clock_settime"
	case 114:
		return "clock_gettime"
	case 115:
		return "clock_getres"
	case 116:
		return "clock_nanosleep"
	case 117:
		return "syslog"
	case 118:
		return "ptrace"
	case 119:
		return "sched_setparam"
	case 120:
		return "sched_setscheduler"
	case 121:
		return "sched_getscheduler"
	case 122:
		return "sched_getparam"
	case 123:
		return "sched_setaffinity"
	case 124:
		return "sched_getaffinity"
	case 125:
		return "sched_yield"
	case 126:
		return "sched_get_priority_max"
	case 127:
		return "sched_get_priority_min"
	case 128:
		return "sched_rr_get_interval"
	case 129:
		return "restart_syscall"
	case 130:
		return "kill"
	case 131:
		return "tkill"
	case 132:
		return "tgkill"
	case 133:
		return "sigaltstack"
	case 134:
		return "rt_sigsuspend"
	case 135:
		return "rt_sigaction"
	case 136:
		return "rt_sigprocmask"
	case 137:
		return "rt_sigpending"
	case 138:
		return "rt_sigtimedwait"
	case 139:
		return "rt_sigqueueinfo"
	case 140:
		return "rt_sigreturn"
	case 141:
		return "setpriority"
	case 142:
		return "getpriority"
	case 143:
		return "reboot"
	case 144:
		return "setregid"
	case 145:
		return "setgid"
	case 146:
		return "setreuid"
	case 147:
		return "setuid"
	case 148:
		return "setresuid"
	case 149:
		return "getresuid"
	case 150:
		return "setresgid"
	case 151:
		return "getresgid"
	case 152:
		return "setfsuid"
	case 153:
		return "setfsgid"
	case 154:
		return "times"
	case 155:
		return "setpgid"
	case 156:
		return "getpgid"
	case 157:
		return "getsid"
	case 158:
		return "setsid"
	case 159:
		return "getgroups"
	case 160:
		return "setgroups"
	case 161:
		return "uname"
	case 162:
		return "sethostname"
	case 163:
		return "setdomainname"
	case 164:
		return "getrlimit"
	case 165:
		return "setrlimit"
	case 166:
		return "getrusage"
	case 167:
		return "umask"
	case 168:
		return "prctl"
	case 169:
		return "getcpu"
	case 170:
		return "gettimeofday"
	case 171:
		return "settimeofday"
	case 172:
		return "adjtimex"
	case 173:
		return "getpid"
	case 174:
		return "getppid"
	case 175:
		return "getuid"
	case 176:
		return "geteuid"
	case 177:
		return "getgid"
	case 178:
		return "getegid"
	case 179:
		return "gettid"
	case 180:
		return "sysinfo"
	case 181:
		return "mq_open"
	case 182:
		return "mq_unlink"
	case 183:
		return "mq_timedsend"
	case 184:
		return "mq_timedreceive"
	case 185:
		return "mq_notify"
	case 186:
		return "mq_getsetattr"
	case 187:
		return "msgget"
	case 188:
		return "msgctl"
	case 189:
		return "msgrcv"
	case 190:
		return "msgsnd"
	case 191:
		return "semget"
	case 192:
		return "semctl"
	case 193:
		return "semtimedop"
	case 194:
		return "semop"
	case 195:
		return "shmget"
	case 196:
		return "shmctl"
	case 197:
		return "shmat"
	case 198:
		return "shmdt"
	case 199:
		return "socket"
	case 200:
		return "socketpair"
	case 201:
		return "bind"
	case 202:
		return "listen"
	case 203:
		return "accept"
	case 204:
		return "connect"
	case 205:
		return "getsockname"
	case 206:
		return "getpeername"
	case 207:
		return "sendto"
	case 208:
		return "recvfrom"
	case 209:
		return "setsockopt"
	case 210:
		return "getsockopt"
	case 211:
		return "shutdown"
	case 212:
		return "sendmsg"
	case 213:
		return "recvmsg"
	case 214:
		return "readahead"
	case 215:
		return "brk"
	case 216:
		return "munmap"
	case 217:
		return "mremap"
	case 218:
		return "add_key"
	case 219:
		return "request_key"
	case 220:
		return "keyctl"
	case 221:
		return "clone"
	case 222:
		return "execve"
	case 223:
		return "mmap"
	case 224:
		return "fadvise64"
	case 225:
		return "swapon"
	case 226:
		return "swapoff"
	case 227:
		return "mprotect"
	case 228:
		return "msync"
	case 229:
		return "mlock"
	case 230:
		return "munlock"
	case 231:
		return "mlockall"
	case 232:
		return "munlockall"
	case 233:
		return "mincore"
	case 234:
		return "madvise"
	case 235:
		return "remap_file_pages"
	case 236:
		return "mbind"
	case 237:
		return "get_mempolicy"
	case 238:
		return "set_mempolicy"
	case 239:
		return "migrate_pages"
	case 240:
		return "move_pages"
	case 241:
		return "rt_tgsigqueueinfo"
	case 242:
		return "perf_event_open"
	case 243:
		return "accept4"
	case 244:
		return "recvmmsg"
	case 245:
		return "arch_specific_syscall"
	case 246:
		return "wait4"
	case 247:
		return "prlimit64"
	case 248:
		return "fanotify_init"
	case 249:
		return "fanotify_mark"
	case 250:
		return "name_to_handle_at"
	case 251:
		return "open_by_handle_at"
	case 252:
		return "clock_adjtime"
	case 253:
		return "syncfs"
	case 254:
		return "setns"
	case 255:
		return "sendmmsg"
	case 256:
		return "process_vm_readv"
	case 257:
		return "process_vm_writev"
	case 258:
		return "kcmp"
	case 259:
		return "finit_module"
	case 260:
		return "sched_setattr"
	case 261:
		return "sched_getattr"
	case 262:
		return "renameat2"
	case 263:
		return "seccomp"
	case 264:
		return "getrandom"
	case 265:
		return "memfd_create"
	case 266:
		return "bpf"
	case 267:
		return "execveat"
	}
	return "unknown"
}

func GetNum(name string) int {
	switch name {
	case "io_setup":
		return 0
	case "io_destroy":
		return 1
	case "io_submit":
		return 2
	case "io_cancel":
		return 3
	case "io_getevents":
		return 4
	case "setxattr":
		return 5
	case "lsetxattr":
		return 6
	case "fsetxattr":
		return 7
	case "getxattr":
		return 8
	case "lgetxattr":
		return 9
	case "fgetxattr":
		return 10
	case "listxattr":
		return 11
	case "llistxattr":
		return 12
	case "flistxattr":
		return 13
	case "removexattr":
		return 14
	case "lremovexattr":
		return 15
	case "fremovexattr":
		return 16
	case "getcwd":
		return 17
	case "lookup_dcookie":
		return 18
	case "eventfd2":
		return 19
	case "epoll_create1":
		return 20
	case "epoll_ctl":
		return 21
	case "epoll_pwait":
		return 22
	case "dup":
		return 23
	case "dup3":
		return 24
	case "fcntl":
		return 25
	case "inotify_init1":
		return 26
	case "inotify_add_watch":
		return 27
	case "inotify_rm_watch":
		return 28
	case "ioctl":
		return 29
	case "ioprio_set":
		return 30
	case "ioprio_get":
		return 31
	case "flock":
		return 32
	case "mknodat":
		return 33
	case "mkdirat":
		return 34
	case "unlinkat":
		return 35
	case "symlinkat":
		return 36
	case "linkat":
		return 37
	case "renameat":
		return 38
	case "umount2":
		return 39
	case "mount":
		return 40
	case "pivot_root":
		return 41
	case "nfsservctl":
		return 42
	case "statfs":
		return 43
	case "fstatfs":
		return 44
	case "truncate":
		return 45
	case "ftruncate":
		return 46
	case "fallocate":
		return 47
	case "faccessat":
		return 48
	case "chdir":
		return 49
	case "fchdir":
		return 50
	case "chroot":
		return 51
	case "fchmod":
		return 52
	case "fchmodat":
		return 53
	case "fchownat":
		return 54
	case "fchown":
		return 55
	case "openat":
		return 56
	case "close":
		return 57
	case "vhangup":
		return 58
	case "pipe2":
		return 59
	case "quotactl":
		return 60
	case "getdents64":
		return 61
	case "lseek":
		return 62
	case "read":
		return 63
	case "write":
		return 64
	case "readv":
		return 65
	case "writev":
		return 66
	case "pread64":
		return 67
	case "pwrite64":
		return 68
	case "preadv":
		return 69
	case "pwritev":
		return 70
	case "sendfile":
		return 71
	case "pselect6":
		return 72
	case "ppoll":
		return 73
	case "signalfd4":
		return 74
	case "vmsplice":
		return 75
	case "splice":
		return 76
	case "tee":
		return 77
	case "readlinkat":
		return 78
	case "fstatat":
		return 79
	case "fstat":
		return 80
	case "sync":
		return 81
	case "fsync":
		return 82
	case "fdatasync":
		return 83
	case "sync_file_range2":
		return 84
	case "sync_file_range":
		return 85
	case "timerfd_create":
		return 86
	case "timerfd_settime":
		return 87
	case "timerfd_gettime":
		return 88
	case "utimensat":
		return 89
	case "acct":
		return 90
	case "capget":
		return 91
	case "capset":
		return 92
	case "personality":
		return 93
	case "exit":
		return 94
	case "exit_group":
		return 95
	case "waitid":
		return 96
	case "set_tid_address":
		return 97
	case "unshare":
		return 98
	case "futex":
		return 99
	case "set_robust_list":
		return 100
	case "get_robust_list":
		return 101
	case "nanosleep":
		return 102
	case "getitimer":
		return 103
	case "setitimer":
		return 104
	case "kexec_load":
		return 105
	case "init_module":
		return 106
	case "delete_module":
		return 107
	case "timer_create":
		return 108
	case "timer_gettime":
		return 109
	case "timer_getoverrun":
		return 110
	case "timer_settime":
		return 111
	case "timer_delete":
		return 112
	case "clock_settime":
		return 113
	case "clock_gettime":
		return 114
	case "clock_getres":
		return 115
	case "clock_nanosleep":
		return 116
	case "syslog":
		return 117
	case "ptrace":
		return 118
	case "sched_setparam":
		return 119
	case "sched_setscheduler":
		return 120
	case "sched_getscheduler":
		return 121
	case "sched_getparam":
		return 122
	case "sched_setaffinity":
		return 123
	case "sched_getaffinity":
		return 124
	case "sched_yield":
		return 125
	case "sched_get_priority_max":
		return 126
	case "sched_get_priority_min":
		return 127
	case "sched_rr_get_interval":
		return 128
	case "restart_syscall":
		return 129
	case "kill":
		return 130
	case "tkill":
		return 131
	case "tgkill":
		return 132
	case "sigaltstack":
		return 133
	case "rt_sigsuspend":
		return 134
	case "rt_sigaction":
		return 135
	case "rt_sigprocmask":
		return 136
	case "rt_sigpending":
		return 137
	case "rt_sigtimedwait":
		return 138
	case "rt_sigqueueinfo":
		return 139
	case "rt_sigreturn":
		return 140
	case "setpriority":
		return 141
	case "getpriority":
		return 142
	case "reboot":
		return 143
	case "setregid":
		return 144
	case "setgid":
		return 145
	case "setreuid":
		return 146
	case "setuid":
		return 147
	case "setresuid":
		return 148
	case "getresuid":
		return 149
	case "setresgid":
		return 150
	case "getresgid":
		return 151
	case "setfsuid":
		return 152
	case "setfsgid":
		return 153
	case "times":
		return 154
	case "setpgid":
		return 155
	case "getpgid":
		return 156
	case "getsid":
		return 157
	case "setsid":
		return 158
	case "getgroups":
		return 159
	case "setgroups":
		return 160
	case "uname":
		return 161
	case "sethostname":
		return 162
	case "setdomainname":
		return 163
	case "getrlimit":
		return 164
	case "setrlimit":
		return 165
	case "getrusage":
		return 166
	case "umask":
		return 167
	case "prctl":
		return 168
	case "getcpu":
		return 169
	case "gettimeofday":
		return 170
	case "settimeofday":
		return 171
	case "adjtimex":
		return 172
	case "getpid":
		return 173
	case "getppid":
		return 174
	case "getuid":
		return 175
	case "geteuid":
		return 176
	case "getgid":
		return 177
	case "getegid":
		return 178
	case "gettid":
		return 179
	case "sysinfo":
		return 180
	case "mq_open":
		return 181
	case "mq_unlink":
		return 182
	case "mq_timedsend":
		return 183
	case "mq_timedreceive":
		return 184
	case "mq_notify":
		return 185
	case "mq_getsetattr":
		return 186
	case "msgget":
		return 187
	case "msgctl":
		return 188
	case "msgrcv":
		return 189
	case "msgsnd":
		return 190
	case "semget":
		return 191
	case "semctl":
		return 192
	case "semtimedop":
		return 193
	case "semop":
		return 194
	case "shmget":
		return 195
	case "shmctl":
		return 196
	case "shmat":
		return 197
	case "shmdt":
		return 198
	case "socket":
		return 199
	case "socketpair":
		return 200
	case "bind":
		return 201
	case "listen":
		return 202
	case "accept":
		return 203
	case "connect":
		return 204
	case "getsockname":
		return 205
	case "getpeername":
		return 206
	case "sendto":
		return 207
	case "recvfrom":
		return 208
	case "setsockopt":
		return 209
	case "getsockopt":
		return 210
	case "shutdown":
		return 211
	case "sendmsg":
		return 212
	case "recvmsg":
		return 213
	case "readahead":
		return 214
	case "brk":
		return 215
	case "munmap":
		return 216
	case "mremap":
		return 217
	case "add_key":
		return 218
	case "request_key":
		return 219
	case "keyctl":
		return 220
	case "clone":
		return 221
	case "execve":
		return 222
	case "mmap":
		return 223
	case "fadvise64":
		return 224
	case "swapon":
		return 225
	case "swapoff":
		return 226
	case "mprotect":
		return 227
	case "msync":
		return 228
	case "mlock":
		return 229
	case "munlock":
		return 230
	case "mlockall":
		return 231
	case "munlockall":
		return 232
	case "mincore":
		return 233
	case "madvise":
		return 234
	case "remap_file_pages":
		return 235
	case "mbind":
		return 236
	case "get_mempolicy":
		return 237
	case "set_mempolicy":
		return 238
	case "migrate_pages":
		return 239
	case "move_pages":
		return 240
	case "rt_tgsigqueueinfo":
		return 241
	case "perf_event_open":
		return 242
	case "accept4":
		return 243
	case "recvmmsg":
		return 244
	case "arch_specific_syscall":
		return 245
	case "wait4":
		return 246
	case "prlimit64":
		return 247
	case "fanotify_init":
		return 248
	case "fanotify_mark":
		return 249
	case "name_to_handle_at":
		return 250
	case "open_by_handle_at":
		return 251
	case "clock_adjtime":
		return 252
	case "syncfs":
		return 253
	case "setns":
		return 254
	case "sendmmsg":
		return 255
	case "process_vm_readv":
		return 256
	case "process_vm_writev":
		return 257
	case "kcmp":
		return 258
	case "finit_module":
		return 259
	case "sched_setattr":
		return 260
	case "sched_getattr":
		return 261
	case "renameat2":
		return 262
	case "seccomp":
		return 263
	case "getrandom":
		return 264
	case "memfd_create":
		return 265
	case "bpf":
		return 266
	case "execveat":
		return 267
	}
	return -1
}

func GetArgsTypes(name string) []ArgType {
	switch name {
	case "io_setup":
		return []ArgType(nil)
	case "io_destroy":
		return []ArgType(nil)
	case "io_submit":
		return []ArgType(nil)
	case "io_cancel":
		return []ArgType(nil)
	case "io_getevents":
		return []ArgType(nil)
	case "setxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR, ARG_INT}
	case "lsetxattr":
		return []ArgType(nil)
	case "fsetxattr":
		return []ArgType(nil)
	case "getxattr":
		return []ArgType{ARG_STR, ARG_STR, ARG_PTR}
	case "lgetxattr":
		return []ArgType(nil)
	case "fgetxattr":
		return []ArgType(nil)
	case "listxattr":
		return []ArgType{ARG_STR, ARG_PTR}
	case "llistxattr":
		return []ArgType(nil)
	case "flistxattr":
		return []ArgType(nil)
	case "removexattr":
		return []ArgType{ARG_STR, ARG_STR}
	case "lremovexattr":
		return []ArgType(nil)
	case "fremovexattr":
		return []ArgType(nil)
	case "getcwd":
		return []ArgType{ARG_PTR}
	case "lookup_dcookie":
		return []ArgType(nil)
	case "eventfd2":
		return []ArgType(nil)
	case "epoll_create1":
		return []ArgType(nil)
	case "epoll_ctl":
		return []ArgType(nil)
	case "epoll_pwait":
		return []ArgType(nil)
	case "dup":
		return []ArgType{ARG_INT}
	case "dup3":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "fcntl":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "inotify_init1":
		return []ArgType(nil)
	case "inotify_add_watch":
		return []ArgType(nil)
	case "inotify_rm_watch":
		return []ArgType(nil)
	case "ioctl":
		return []ArgType(nil)
	case "ioprio_set":
		return []ArgType(nil)
	case "ioprio_get":
		return []ArgType(nil)
	case "flock":
		return []ArgType{ARG_INT, ARG_INT}
	case "mknodat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "mkdirat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT}
	case "unlinkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT}
	case "symlinkat":
		return []ArgType{ARG_STR, ARG_INT, ARG_STR}
	case "linkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_STR, ARG_INT}
	case "renameat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_STR}
	case "umount2":
		return []ArgType(nil)
	case "mount":
		return []ArgType{ARG_STR, ARG_STR, ARG_STR, ARG_INT, ARG_PTR}
	case "pivot_root":
		return []ArgType(nil)
	case "nfsservctl":
		return []ArgType(nil)
	case "statfs":
		return []ArgType{ARG_STR, ARG_PTR}
	case "fstatfs":
		return []ArgType{ARG_INT, ARG_PTR}
	case "truncate":
		return []ArgType{ARG_STR, ARG_PTR}
	case "ftruncate":
		return []ArgType{ARG_INT, ARG_PTR}
	case "fallocate":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "faccessat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "chdir":
		return []ArgType{ARG_STR}
	case "fchdir":
		return []ArgType{ARG_INT}
	case "chroot":
		return []ArgType{ARG_STR}
	case "fchmod":
		return []ArgType{ARG_INT, ARG_INT}
	case "fchmodat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "fchownat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT, ARG_INT}
	case "fchown":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "openat":
		return []ArgType{ARG_INT, ARG_STR, ARG_INT, ARG_INT}
	case "close":
		return []ArgType{ARG_INT}
	case "vhangup":
		return []ArgType(nil)
	case "pipe2":
		return []ArgType{ARG_PTR, ARG_INT}
	case "quotactl":
		return []ArgType(nil)
	case "getdents64":
		return []ArgType(nil)
	case "lseek":
		return []ArgType(nil)
	case "read":
		return []ArgType{ARG_INT, ARG_PTR}
	case "write":
		return []ArgType{ARG_INT, ARG_PTR}
	case "readv":
		return []ArgType(nil)
	case "writev":
		return []ArgType(nil)
	case "pread64":
		return []ArgType(nil)
	case "pwrite64":
		return []ArgType(nil)
	case "preadv":
		return []ArgType(nil)
	case "pwritev":
		return []ArgType(nil)
	case "sendfile":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "pselect6":
		return []ArgType(nil)
	case "ppoll":
		return []ArgType(nil)
	case "signalfd4":
		return []ArgType(nil)
	case "vmsplice":
		return []ArgType(nil)
	case "splice":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_INT, ARG_INT}
	case "tee":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "readlinkat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "fstatat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR, ARG_INT}
	case "fstat":
		return []ArgType{ARG_INT, ARG_PTR}
	case "sync":
		return []ArgType(nil)
	case "fsync":
		return []ArgType{ARG_INT}
	case "fdatasync":
		return []ArgType{ARG_INT}
	case "sync_file_range2":
		return []ArgType(nil)
	case "sync_file_range":
		return []ArgType(nil)
	case "timerfd_create":
		return []ArgType(nil)
	case "timerfd_settime":
		return []ArgType(nil)
	case "timerfd_gettime":
		return []ArgType(nil)
	case "utimensat":
		return []ArgType{ARG_INT, ARG_STR, ARG_PTR}
	case "acct":
		return []ArgType{ARG_STR}
	case "capget":
		return []ArgType(nil)
	case "capset":
		return []ArgType(nil)
	case "personality":
		return []ArgType(nil)
	case "exit":
		return []ArgType{ARG_INT}
	case "exit_group":
		return []ArgType(nil)
	case "waitid":
		return []ArgType(nil)
	case "set_tid_address":
		return []ArgType(nil)
	case "unshare":
		return []ArgType{ARG_INT}
	case "futex":
		return []ArgType(nil)
	case "set_robust_list":
		return []ArgType(nil)
	case "get_robust_list":
		return []ArgType(nil)
	case "nanosleep":
		return []ArgType{ARG_PTR, ARG_PTR}
	case "getitimer":
		return []ArgType(nil)
	case "setitimer":
		return []ArgType(nil)
	case "kexec_load":
		return []ArgType(nil)
	case "init_module":
		return []ArgType(nil)
	case "delete_module":
		return []ArgType(nil)
	case "timer_create":
		return []ArgType(nil)
	case "timer_gettime":
		return []ArgType(nil)
	case "timer_getoverrun":
		return []ArgType(nil)
	case "timer_settime":
		return []ArgType(nil)
	case "timer_delete":
		return []ArgType(nil)
	case "clock_settime":
		return []ArgType(nil)
	case "clock_gettime":
		return []ArgType(nil)
	case "clock_getres":
		return []ArgType(nil)
	case "clock_nanosleep":
		return []ArgType(nil)
	case "syslog":
		return []ArgType(nil)
	case "ptrace":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT}
	case "sched_setparam":
		return []ArgType(nil)
	case "sched_setscheduler":
		return []ArgType(nil)
	case "sched_getscheduler":
		return []ArgType(nil)
	case "sched_getparam":
		return []ArgType(nil)
	case "sched_setaffinity":
		return []ArgType(nil)
	case "sched_getaffinity":
		return []ArgType(nil)
	case "sched_yield":
		return []ArgType(nil)
	case "sched_get_priority_max":
		return []ArgType(nil)
	case "sched_get_priority_min":
		return []ArgType(nil)
	case "sched_rr_get_interval":
		return []ArgType(nil)
	case "restart_syscall":
		return []ArgType(nil)
	case "kill":
		return []ArgType{ARG_INT, ARG_PTR}
	case "tkill":
		return []ArgType(nil)
	case "tgkill":
		return []ArgType{ARG_INT, ARG_INT, ARG_PTR}
	case "sigaltstack":
		return []ArgType(nil)
	case "rt_sigsuspend":
		return []ArgType(nil)
	case "rt_sigaction":
		return []ArgType(nil)
	case "rt_sigprocmask":
		return []ArgType(nil)
	case "rt_sigpending":
		return []ArgType(nil)
	case "rt_sigtimedwait":
		return []ArgType(nil)
	case "rt_sigqueueinfo":
		return []ArgType(nil)
	case "rt_sigreturn":
		return []ArgType(nil)
	case "setpriority":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getpriority":
		return []ArgType{ARG_INT, ARG_INT}
	case "reboot":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_STR}
	case "setregid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setgid":
		return []ArgType{ARG_INT}
	case "setreuid":
		return []ArgType{ARG_INT, ARG_INT}
	case "setuid":
		return []ArgType{ARG_INT}
	case "setresuid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getresuid":
		return []ArgType(nil)
	case "setresgid":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "getresgid":
		return []ArgType(nil)
	case "setfsuid":
		return []ArgType{ARG_INT}
	case "setfsgid":
		return []ArgType{ARG_INT}
	case "times":
		return []ArgType{ARG_PTR}
	case "setpgid":
		return []ArgType{ARG_INT, ARG_INT}
	case "getpgid":
		return []ArgType{ARG_INT}
	case "getsid":
		return []ArgType(nil)
	case "setsid":
		return []ArgType(nil)
	case "getgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setgroups":
		return []ArgType{ARG_INT, ARG_PTR}
	case "uname":
		return []ArgType{ARG_PTR}
	case "sethostname":
		return []ArgType{ARG_PTR}
	case "setdomainname":
		return []ArgType{ARG_PTR}
	case "getrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "setrlimit":
		return []ArgType{ARG_INT, ARG_PTR}
	case "getrusage":
		return []ArgType{ARG_INT, ARG_PTR}
	case "umask":
		return []ArgType{ARG_INT}
	case "prctl":
		return []ArgType(nil)
	case "getcpu":
		return []ArgType(nil)
	case "gettimeofday":
		return []ArgType{ARG_PTR}
	case "settimeofday":
		return []ArgType{ARG_PTR}
	case "adjtimex":
		return []ArgType{ARG_PTR}
	case "getpid":
		return []ArgType(nil)
	case "getppid":
		return []ArgType(nil)
	case "getuid":
		return []ArgType(nil)
	case "geteuid":
		return []ArgType(nil)
	case "getgid":
		return []ArgType(nil)
	case "getegid":
		return []ArgType(nil)
	case "gettid":
		return []ArgType(nil)
	case "sysinfo":
		return []ArgType{ARG_PTR}
	case "mq_open":
		return []ArgType(nil)
	case "mq_unlink":
		return []ArgType(nil)
	case "mq_timedsend":
		return []ArgType(nil)
	case "mq_timedreceive":
		return []ArgType(nil)
	case "mq_notify":
		return []ArgType(nil)
	case "mq_getsetattr":
		return []ArgType(nil)
	case "msgget":
		return []ArgType(nil)
	case "msgctl":
		return []ArgType(nil)
	case "msgrcv":
		return []ArgType(nil)
	case "msgsnd":
		return []ArgType(nil)
	case "semget":
		return []ArgType(nil)
	case "semctl":
		return []ArgType(nil)
	case "semtimedop":
		return []ArgType(nil)
	case "semop":
		return []ArgType(nil)
	case "shmget":
		return []ArgType(nil)
	case "shmctl":
		return []ArgType(nil)
	case "shmat":
		return []ArgType(nil)
	case "shmdt":
		return []ArgType(nil)
	case "socket":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT}
	case "socketpair":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "bind":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "listen":
		return []ArgType{ARG_INT, ARG_INT}
	case "accept":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "connect":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getsockname":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "getpeername":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR}
	case "sendto":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "recvfrom":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR, ARG_PTR}
	case "setsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_INT}
	case "getsockopt":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_PTR, ARG_PTR}
	case "shutdown":
		return []ArgType{ARG_INT, ARG_INT}
	case "sendmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "recvmsg":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT}
	case "readahead":
		return []ArgType(nil)
	case "brk":
		return []ArgType(nil)
	case "munmap":
		return []ArgType{ARG_INT, ARG_INT}
	case "mremap":
		return []ArgType(nil)
	case "add_key":
		return []ArgType(nil)
	case "request_key":
		return []ArgType(nil)
	case "keyctl":
		return []ArgType(nil)
	case "clone":
		return []ArgType(nil)
	case "execve":
		return []ArgType(nil)
	case "mmap":
		return []ArgType{ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_INT, ARG_PTR}
	case "fadvise64":
		return []ArgType(nil)
	case "swapon":
		return []ArgType(nil)
	case "swapoff":
		return []ArgType(nil)
	case "mprotect":
		return []ArgType{ARG_PTR, ARG_INT}
	case "msync":
		return []ArgType(nil)
	case "mlock":
		return []ArgType{ARG_PTR}
	case "munlock":
		return []ArgType{ARG_PTR}
	case "mlockall":
		return []ArgType{ARG_INT}
	case "munlockall":
		return []ArgType(nil)
	case "mincore":
		return []ArgType(nil)
	case "madvise":
		return []ArgType{ARG_PTR, ARG_INT}
	case "remap_file_pages":
		return []ArgType(nil)
	case "mbind":
		return []ArgType(nil)
	case "get_mempolicy":
		return []ArgType(nil)
	case "set_mempolicy":
		return []ArgType(nil)
	case "migrate_pages":
		return []ArgType(nil)
	case "move_pages":
		return []ArgType(nil)
	case "rt_tgsigqueueinfo":
		return []ArgType(nil)
	case "perf_event_open":
		return []ArgType(nil)
	case "accept4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_PTR, ARG_INT}
	case "recvmmsg":
		return []ArgType(nil)
	case "arch_specific_syscall":
		return []ArgType(nil)
	case "wait4":
		return []ArgType{ARG_INT, ARG_PTR, ARG_INT, ARG_PTR}
	case "prlimit64":
		return []ArgType(nil)
	case "fanotify_init":
		return []ArgType(nil)
	case "fanotify_mark":
		return []ArgType(nil)
	case "name_to_handle_at":
		return []ArgType(nil)
	case "open_by_handle_at":
		return []ArgType(nil)
	case "clock_adjtime":
		return []ArgType(nil)
	case "syncfs":
		return []ArgType(nil)
	case "setns":
		return []ArgType(nil)
	case "sendmmsg":
		return []ArgType(nil)
	case "process_vm_readv":
		return []ArgType(nil)
	case "process_vm_writev":
		return []ArgType(nil)
	case "kcmp":
		return []ArgType(nil)
	case "finit_module":
		return []ArgType(nil)
	case "sched_setattr":
		return []ArgType(nil)
	case "sched_getattr":
		return []ArgType(nil)
	case "renameat2":
		return []ArgType(nil)
	case "seccomp":
		return []ArgType(nil)
	case "getrandom":
		return []ArgType(nil)
	case "memfd_create":
		return []ArgType(nil)
	case "bpf":
		return []ArgType(nil)
	case "execveat":
		return []ArgType(nil)
	}
	return nil
}
