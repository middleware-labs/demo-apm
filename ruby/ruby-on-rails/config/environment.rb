# Load the Rails application.
require_relative "application"

require 'middleware/ruby_gem'

Middleware::RubyGem.init

# Initialize the Rails application.
Rails.application.initialize!
