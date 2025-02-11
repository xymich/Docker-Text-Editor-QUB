FROM ruby:3.3.6-alpine

RUN apk add --update --no-cache build-base libxml2-dev libxslt-dev

WORKDIR /usr/src/app

COPY Gemfile Gemfile.lock ./

RUN bundle install

COPY . .

EXPOSE 8020

CMD ["bundle", "exec", "rackup", "--host", "0.0.0.0", "--port", "8020"]
