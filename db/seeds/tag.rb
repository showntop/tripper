#!/usr/bin/env ruby
require 'pp'

class Tag

	def result
		@tag_hash
	end

	def initialize(version, path, conn, json_obj_array)
		@conn = conn
		@path = path
		@version = version
		@json_obj_array = json_obj_array

		@tag_hash = {}
	end

	def seed
		seed_sql @json_obj_array.map { |e| e["tages"].split('$') }.flatten.uniq
	end

	private
	def seed_sql tags
		#find the biggest id of tag
		base_id = "#{Time.now.strftime('%y%m')}0000".to_i
		file = File.open(@path + "/tmp_tag_sql_" + @version, "w") 
		tags.each_with_index do |c, i|
			@tag_hash[c] = base_id + i
			file.puts "insert into tags(id, name, created_at, updated_at) values(#{base_id + i}, '#{c}', '#{Time.now.strftime('%Y-%m-%d %H:%M:%S')}', '#{Time.now.strftime('%Y-%m-%d %H:%M:%S')}');"
		end
		file.close
	end

end