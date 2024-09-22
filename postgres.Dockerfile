FROM postgres:14-alpine

COPY ./src/init.sql /tmp/init.sql
RUN #psql postgres < /tmp/init.sql