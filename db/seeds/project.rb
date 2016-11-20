#!/usr/bin/env ruby
require 'pp'

class Project

	def initialize(version, path, conn, json_obj_array, category_hash, tag_hash)
		@conn = conn
		@path = path
		@version = version
		@json_obj_array = json_obj_array
		@category_hash = category_hash
		@tag_hash = tag_hash
	end

	def seed
		base_id = "#{Time.now.strftime('%y%m')}0000".to_i
		project_array = Array.new
		@json_obj_array.each_with_index do |e, i|
 			item = Hash.new
 			item["id"] = base_id + i
 			item["name"] = e["name"]
 			item["description"] = e["info"].gsub(/\r\n/, "<br>").gsub("'","''")
 			item["version"] = e["version"]
 			item["size"] = e["size"]
 			item["dlink"] = e["download_url"]
 			item["logo_url"] = e["icon_url"]
 			item["category_id"] = @category_hash[e["type"]]
 			create_tags((base_id + i), e['tages'].split('$'))
 			project_array.push item
		end
		seed_sql project_array
	end

	private
	def seed_sql projects
		#find the biggest id of project
		file = File.open(@path + "/tmp_project_sql_" + @version, "w") 
		projects.each do |p|
			file.puts "insert into projects(#{p.keys.join(',')}) values(#{p['id']},'#{p['name']}','#{p['description']}', '#{p['version']}', #{(p['size'].to_f * 1024 * 1024)}, '#{p['dlink']}', '#{p['logo_url']}', #{p['category_id']});"
		end
		file.close
	end



	def create_tags(pid, tags)
		file = File.new(@path + "/tmp_project_tag_sql_" + @version, "a") 
		tags.each do |t|
			tag_id = @tag_hash[t]
			file.puts "insert into project_tags(project_id, tag_id) values(#{pid}, #{tag_id});"
		end
		file.close
	end

end


