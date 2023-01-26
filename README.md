# Woof
[woof](http://www.home.unix-ag.org/simon/woof.html) is a fantastic utility to transfer files between two PC's in the same LAN over HTTP.  However, it requires a bulky Python runtime on the host machine.  This project is a woof-clone but written in go to avoid this restriction.

# Installation
```bash
git clone https://github.com/tvanriel/woof.git
make all
make install
```
---
`woof` is now installed for your local user.