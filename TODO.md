## Todos in this project

- [x] Reads `--dir` flag and uses this as the project directory. (Medium)
- [ ] Dynamically add attributes to goland config file. (Medium)
- [ ] Handle orphan docker processes. (Hard)
- [ ] Better error handling. Ex: if the docker build or docker run fails. (Medium)
- [ ] Set a max replicas limit. (Easy)
- [x] Support for shorthand flags (Easy)
- [ ] Ideally append to the `globals.RUNNING_IDS` in the `RunDocker` function. This would handle case where a `SIGTERM` is triggered while the docker processes are being spawned, so the cleanup will not catch the previously spawned processes because the `globals.RUNNING_IDS` are appended after all the docker processes have been spawned. This would require go routine synchronization with a mutex lock. (Medium)
