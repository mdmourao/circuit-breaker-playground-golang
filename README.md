# circuit-breaker-playground-golang

```
2025/10/02 08:32:49 Starting client
2025/10/02 08:32:50 Request succeeded after 1.095177ms: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:51 Request succeeded after 471.011µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:52 Request succeeded after 347.565µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:53 Request succeeded after 375.013µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:54 Request succeeded after 363.64µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:55 Request succeeded after 360.514µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:56 Request succeeded after 359.9µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:57 Request succeeded after 380.967µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:58 Request succeeded after 345.948µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:32:59 Request succeeded after 364.111µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:00 Request succeeded after 376.089µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:01 Request succeeded after 340.929µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:02 Request succeeded after 382.968µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:03 Request succeeded after 349.177µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:04 Request succeeded after 543.563µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:05 Request succeeded after 370.178µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:06 Request succeeded after 407.574µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:07 Request succeeded after 312.457µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:33:08 Request succeeded after 394.4µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
// stopped server
2025/10/02 08:33:09 Request failed after 512.981µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:10 Request failed after 524.573µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:11 Request failed after 480.042µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:12 Request failed after 518.236µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:13 Request failed after 1.594787ms: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:14 Request failed after 511.288µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:15 Request failed after 499.659µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:16 Request failed after 528.579µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:17 Request failed after 457.529µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:18 Request failed after 391.355µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:19 Request failed after 495.214µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:20 Request failed after 436.09µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:21 Request failed after 469.83µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:22 Request failed after 423.413µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:23 Request failed after 351.282µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:24 Request failed after 469.934µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:25 Request failed after 431.691µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:26 Request failed after 462.981µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:27 Request failed after 616.341µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:28 Request failed after 432.741µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:29 Request failed after 443.735µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:30 Request failed after 456.045µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:31 Request failed after 459.125µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:32 Request failed after 471.286µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:33 Request failed after 482.64µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:34 Request failed after 371.561µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:35 Request failed after 457.674µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:36 Request failed after 1.723441ms: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:37 Circuit breaker 'BladeRunnerAPI' changed from closed to open
2025/10/02 08:33:37 Request failed after 562.043µs: Get "http://localhost:8080/blade-runner": dial tcp [::1]:8080: connect: connection refused
2025/10/02 08:33:38 Request failed after 1.267µs: circuit breaker is open
2025/10/02 08:33:39 Request failed after 1.595µs: circuit breaker is open
2025/10/02 08:33:40 Request failed after 1.417µs: circuit breaker is open
2025/10/02 08:33:41 Request failed after 1.277µs: circuit breaker is open
2025/10/02 08:33:42 Request failed after 1.117µs: circuit breaker is open
2025/10/02 08:33:43 Request failed after 2.152µs: circuit breaker is open
2025/10/02 08:33:44 Request failed after 1.176µs: circuit breaker is open
2025/10/02 08:33:45 Request failed after 1.072µs: circuit breaker is open
2025/10/02 08:33:46 Request failed after 1.132µs: circuit breaker is open
2025/10/02 08:33:47 Request failed after 1.283µs: circuit breaker is open
2025/10/02 08:33:48 Request failed after 1.51µs: circuit breaker is open
2025/10/02 08:33:49 Request failed after 1.97µs: circuit breaker is open
2025/10/02 08:33:50 Request failed after 1.368µs: circuit breaker is open
2025/10/02 08:33:51 Request failed after 1.23µs: circuit breaker is open
2025/10/02 08:33:52 Request failed after 1.126µs: circuit breaker is open
2025/10/02 08:33:53 Request failed after 1.18µs: circuit breaker is open
2025/10/02 08:33:54 Request failed after 1.768µs: circuit breaker is open
2025/10/02 08:33:55 Request failed after 1.142µs: circuit breaker is open
2025/10/02 08:33:56 Request failed after 1.219µs: circuit breaker is open
2025/10/02 08:33:57 Request failed after 1.236µs: circuit breaker is open
2025/10/02 08:33:58 Request failed after 968ns: circuit breaker is open
2025/10/02 08:33:59 Request failed after 1.181µs: circuit breaker is open
2025/10/02 08:34:00 Request failed after 1.008µs: circuit breaker is open
2025/10/02 08:34:01 Request failed after 1.183µs: circuit breaker is open
2025/10/02 08:34:02 Request failed after 1.914µs: circuit breaker is open
2025/10/02 08:34:03 Request failed after 1.197µs: circuit breaker is open
2025/10/02 08:34:04 Request failed after 1.621µs: circuit breaker is open
2025/10/02 08:34:05 Request failed after 1.17µs: circuit breaker is open
2025/10/02 08:34:06 Request failed after 1.012µs: circuit breaker is open
2025/10/02 08:34:07 Request failed after 1.32µs: circuit breaker is open
2025/10/02 08:34:08 Request failed after 1.295µs: circuit breaker is open
2025/10/02 08:34:09 Request failed after 1.149µs: circuit breaker is open
2025/10/02 08:34:10 Request failed after 1.088µs: circuit breaker is open
2025/10/02 08:34:11 Request failed after 1.117µs: circuit breaker is open
2025/10/02 08:34:12 Request failed after 1.119µs: circuit breaker is open
2025/10/02 08:34:13 Request failed after 1.341µs: circuit breaker is open
2025/10/02 08:34:14 Request failed after 1.042µs: circuit breaker is open
2025/10/02 08:34:15 Request failed after 1.535µs: circuit breaker is open
2025/10/02 08:34:16 Request failed after 1.259µs: circuit breaker is open
2025/10/02 08:34:17 Request failed after 1.821µs: circuit breaker is open
2025/10/02 08:34:18 Request failed after 1.33µs: circuit breaker is open
2025/10/02 08:34:19 Request failed after 1.117µs: circuit breaker is open
2025/10/02 08:34:20 Request failed after 1.26µs: circuit breaker is open
2025/10/02 08:34:21 Request failed after 1.331µs: circuit breaker is open
// start server
2025/10/02 08:34:22 Request failed after 1.762µs: circuit breaker is open
2025/10/02 08:34:23 Request failed after 1.24µs: circuit breaker is open
2025/10/02 08:34:24 Request failed after 1.303µs: circuit breaker is open
2025/10/02 08:34:25 Request failed after 1.105µs: circuit breaker is open
2025/10/02 08:34:26 Request failed after 1.12µs: circuit breaker is open
2025/10/02 08:34:27 Request failed after 1.124µs: circuit breaker is open
2025/10/02 08:34:28 Request failed after 1.195µs: circuit breaker is open
2025/10/02 08:34:29 Request failed after 1.414µs: circuit breaker is open
2025/10/02 08:34:30 Request failed after 1.159µs: circuit breaker is open
2025/10/02 08:34:31 Request failed after 1.16µs: circuit breaker is open
2025/10/02 08:34:32 Request failed after 1.133µs: circuit breaker is open
2025/10/02 08:34:33 Request failed after 1.138µs: circuit breaker is open
2025/10/02 08:34:34 Request failed after 1.172µs: circuit breaker is open
2025/10/02 08:34:35 Request failed after 1.189µs: circuit breaker is open
2025/10/02 08:34:36 Request failed after 1.382µs: circuit breaker is open
2025/10/02 08:34:37 Request failed after 1.468µs: circuit breaker is open
2025/10/02 08:34:38 Circuit breaker 'BladeRunnerAPI' changed from open to half-open
2025/10/02 08:34:38 Request succeeded after 889.787µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:39 Request succeeded after 804.405µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:40 Circuit breaker 'BladeRunnerAPI' changed from half-open to closed
2025/10/02 08:34:40 Request succeeded after 390.858µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:41 Request succeeded after 354.681µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:42 Request succeeded after 373.853µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:43 Request succeeded after 370.241µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:44 Request succeeded after 360.235µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:45 Request succeeded after 368.847µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:46 Request succeeded after 359.218µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:47 Request succeeded after 334.486µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:48 Request succeeded after 580.419µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
2025/10/02 08:34:49 Request succeeded after 361.797µs: {"message":"I've seen things you people wouldn't believe...","movie":"Blade Runner","year":1982}
```
