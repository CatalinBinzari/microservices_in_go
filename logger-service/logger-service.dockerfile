
FROM alpine:latest

RUN mkdir /app
# RUN mkdir -p /data/db
# RUN chmod -R go+w /data/db

COPY  loggerServiceApp /app

CMD [ "/app/loggerServiceApp" ]