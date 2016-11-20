#!/usr/bin/env ruby
require 'rubygems'
require 'multi_json'
require 'pp'
require 'pg'

require './category'
require './project'
require './tag'

DATA_VERSION = '20161109'
DATA_PATH = '../data'

# Output a table of current connections to the DB
# $conn = PG.connect(host: 'ec2-54-243-201-19.compute-1.amazonaws.com', dbname: 'dbl153jnanfv0h', user:'ynazkzqhiihmle', password: 'sLlrp2_gLeJCoGYvVXt3Ts1jNY' )
pp '1.connect postgres success'

full_file_path = DATA_PATH + '/' +  'records-' + DATA_VERSION + '.json'
pp '2.load'+full_file_path
json_file = File.read(full_file_path)
json_obj = MultiJson.load(json_file)
# pp json_obj

category = Category.new(DATA_VERSION, DATA_PATH, $conn, json_obj)
category.seed
pp '3.done category'

tag = Tag.new(DATA_VERSION, DATA_PATH, $conn, json_obj)
tag.seed
pp '4.done tag'

project = Project.new(DATA_VERSION, DATA_PATH, $conn, json_obj, category.result, tag.result)
project.seed
pp '5.done project'
