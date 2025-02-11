require 'spec_helper'
require 'rack/test'
require 'json'
require_relative '../app'

ENV['RACK_ENV'] = 'test'

RSpec.describe 'My Sinatra Application' do
  include Rack::Test::Methods

  def app
    Sinatra::Application
  end

  describe 'GET /reverse' do
    it 'reverses the text' do
      get '/reverse', text: 'Hello, World!'
      expect(last_response.status).to eq(200)
      response = JSON.parse(last_response.body)
      expect(response['error']).to eq(false)
      expect(response['string']).to eq('Reversed text')
      expect(response['answer']).to eq('!dlroW ,olleH')
    end

    it 'returns an error for empty text' do
      get '/reverse', text: ''
      expect(last_response.status).to eq(400)
      response = JSON.parse(last_response.body)
      expect(response['error']).to eq(true)
      expect(response['string']).to eq('Missing "text" parameter')
      expect(response['answer']).to eq('')
    end
  end
end
