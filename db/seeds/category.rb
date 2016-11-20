#!/usr/bin/env ruby
require 'pp'

class Category

	def result
		@category_hash
	end

	def initialize(version, path, conn, json_obj_array)
		@conn = conn
		@path = path
		@version = version
		@json_obj_array = json_obj_array

		@category_hash = {}
	end

	def seed
		seed_sql @json_obj_array.map { |e| e["type"] }.uniq
	end

	private
	def seed_sql categories
		#find the biggest id of category
		base_id = "#{Time.now.strftime('%y%m')}0000".to_i
		file = File.open(@path + "/tmp_category_sql_" + @version, "w") 
		categories.each_with_index do |c, i|
			@category_hash[c] = base_id + i
			file.puts "insert into categories(id, name, created_at, updated_at) values(#{base_id + i}, '#{c}', '#{Time.now.strftime('%Y-%m-%d %H:%M:%S')}', '#{Time.now.strftime('%Y-%m-%d %H:%M:%S')}');"
		end
		file.close
	end

end