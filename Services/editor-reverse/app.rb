require 'sinatra'
require 'json'

set :port, ENV['PORT'] || 8020

get '/reverse' do
  content_type :json
  text = params['text']
  if text.nil? || text.empty?
    status 400
    return { error: true, string: 'Missing "text" parameter', answer: '' }.to_json
  end

  reversed_text = text.reverse
  { error: false, string: 'Reversed text', answer: reversed_text }.to_json
end

# Start the server if this file is executed directly
if __FILE__ == $0
  Sinatra::Application.run!
end
