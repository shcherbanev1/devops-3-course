Так как выполняю задания с WSL и у меня нет сервака на Linux а поднимать еще одну WSL не хочу
То поднял docker контейнер с ubuntu

В контейнер установил ssh сервер
Перекинул публичный ключ с WSL в контейнер

создаю inventory.ini где описал адрес ssh сервера
и пингую:
kostya@konstantin:~/devops-3-course/2task$ ansible -i inventory.ini servers -m ping
[WARNING]: Platform linux on host test is using the discovered Python interpreter at /usr/bin/python3.12, but future installation of another Python interpreter could change the meaning of that path.
See https://docs.ansible.com/ansible-core/2.18/reference_appendices/interpreter_discovery.html for more information.
test | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python3.12"
    },
    "changed": false,
    "ping": "pong"
}
kostya@konstantin:~/devops-3-course/2task$ 

написал плейбук
запустил его - задачи выполнились

Проверяю:
root@c0702a5bf13f:~# id student
uid=1001(student) gid=1001(student) groups=1001(student),27(sudo)

и подключившись по ssh от стьюдента
kostya@konstantin:~/devops-3-course/2task$ ssh -i ~/.ssh/id_ed25519 -p 2222 student@127.0.0.1
Welcome to Ubuntu 24.04.2 LTS (GNU/Linux 5.15.167.4-microsoft-standard-WSL2 x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/pro

This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Tue Sep 15 20:53:51 2026 from 172.17.0.1
To run a command as administrator (user "root"), use "sudo <command>".
See "man sudo_root" for details.

student@c0702a5bf13f:~$ 