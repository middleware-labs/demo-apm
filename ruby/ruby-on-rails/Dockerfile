FROM ruby:3.0.2

WORKDIR /app

COPY Gemfile Gemfile.lock ./

RUN bundle install

COPY . .

EXPOSE 3000

CMD ["bin/rails", "server"]
