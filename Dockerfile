FROM baseimage-go

WORKDIR /

COPY . .

EXPOSE 9000

ENTRYPOINT [ "./compileandrunserver.sh" ]