FROM ruby:3.0.2

WORKDIR /app

COPY Gemfile Gemfile.lock ./

RUN bundle install

COPY . .

RUN bundle exec rails db:migrate

CMD ["bin/rails", "server", "-b", "0.0.0.0",  "-p", "3000"]
