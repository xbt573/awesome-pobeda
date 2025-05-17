#!/usr/bin/env ruby
require 'json'
require 'net/http'
require 'erb'
require 'optparse'
require_relative 'bundled/simpleidn'

options = {
    output: "output"
}

OptionParser.new do |parser|
    parser.on('-i', '--input=FILE', 'Input file (JSON)')
    parser.on('-a', '--availability=FILE', 'Availability file (JSON)')
    parser.on('-t', '--template=FILE', 'Template file (ERB)')
    parser.on('-o', '--output=FILE', 'Output file')
end.parse!(into: options)

[:input, :availability, :template].each do |option|
    fail "Missing required option: #{option}" unless options[option] != nil
    fail "File \"#{options[option]}\" doesn't exist or is directory" unless File.file?(options[option])
end

data = JSON.parse(File.read(options[:input]))
availability = JSON.parse(File.read(options[:availability]))

template = ERB.new(File.read(options[:template]), trim_mode: '-')

File.write options[:output], template.result
