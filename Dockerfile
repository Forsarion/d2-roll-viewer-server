FROM swift as environment

RUN apt-get update
RUN apt-get install uuid-dev

EXPOSE 8181

WORKDIR /app

FROM environment

COPY . .

RUN swift build

CMD swift run