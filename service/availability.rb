#!/usr/bin/env ruby
require 'json'
require 'net/http'
require 'optparse'
require_relative 'bundled/simpleidn'

options = {
  output: "availability.json"
}

OptionParser.new do |parser|
    parser.on('-i', '--input=FILE', 'Input file (JSON)')
    parser.on('-o', '--output=FILE', 'Output file (JSON)')
end.parse!(into: options)

[:input].each do |option|
    fail "Missing required option: #{option}" unless options[option] != nil
    fail "File \"#{options[option]}\" doesn't exist or is directory" unless File.file?(options[option])
end

data = JSON.parse(File.read(options[:input]))

availability = {}

data.each do |category|
    category["items"] = category["items"].sort.to_h

    category["items"].each do |key, domains|
        domains.each do |domain|
            begin
                Net::HTTP.get(URI("https://#{SimpleIDN.to_ascii(domain)}"))
            rescue OpenSSL::SSL::SSLError
                begin
                    Net::HTTP.get(URI("http://#{SimpleIDN.to_ascii(domain)}"))
                rescue
                    availability[domain] = [false]
                else
                    availability[domain] = [true, "http"]
                end
            else
                availability[domain] = [true, "https"]
            end
        end
    end
end

File.write options[:output], JSON.generate(availability)
