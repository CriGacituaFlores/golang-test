class User < ApplicationRecord
  include HTTParty

  def self.get_all_filtered
    a = HTTParty.get("http://localhost:3001/get_users_filtered")
    return a
  end
end
