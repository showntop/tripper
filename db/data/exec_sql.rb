#!/usr/bin/env ruby
require 'rubygems'
require 'pp'
require 'pg'

DATA_VERSION = '20161109'

tables = ['category', 'tag', 'project', 'project_tag']

$conn = PG.connect(host: 'ec2-54-243-201-19.compute-1.amazonaws.com', dbname: 'dbl153jnanfv0h', user:'ynazkzqhiihmle', password: 'sLlrp2_gLeJCoGYvVXt3Ts1jNY' )
pp '1.connect postgres success'
pp '2.migrate data'
tables.each do |table|
	filename = "tmp_#{table}_sql_#{DATA_VERSION}"
	pp "do ...#{filename}"
	IO.foreach(filename) do |block| 
		begin
			$conn.exec(block)
		rescue Exception => e
		  	pp e
		end
	end
end
pp '3.done'
