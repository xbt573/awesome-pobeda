#!/usr/bin/env ruby
require 'json'
require 'erb'
require 'optparse'

options = {
    output: "output"
}

OptionParser.new do |parser|
    parser.on('-i', '--input=FILE', 'Input file (JSON)')
    parser.on('-t', '--template=FILE', 'Template file (ERB)')
    parser.on('-o', '--output=FILE', 'Output file')
end.parse!(into: options)

[:input, :template].each do |option|
    fail "Missing required option: #{option}" unless options[option] != nil
    fail "File \"#{options[option]}\" doesn't exist or is directory" unless File.file?(options[option])
end

data = JSON.parse(File.read(options[:input]))
template = ERB.new(File.read(options[:template]), trim_mode: '-')

File.write options[:output], template.result
