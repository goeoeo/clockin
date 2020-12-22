FROM centos

WORKDIR "/usr/local/npc"

ADD ["npc","npc"]

RUN ["chmod","+x","npc"]

CMD ["./npc","-server=cchuiyu.cn:8024","-vkey=h2j9a9xsjyjsnsp0","-type=tcp"]