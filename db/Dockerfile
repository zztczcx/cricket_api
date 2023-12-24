FROM migrate/migrate

COPY ./migrations /migrations

ENTRYPOINT [ "migrate", "-path", "/migrations", "-database"]
CMD ["mysql://root:password@tcp(cricket.db:3306)/cricket_db up"]
