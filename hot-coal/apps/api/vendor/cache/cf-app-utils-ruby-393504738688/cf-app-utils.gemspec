# -*- encoding: utf-8 -*-
# stub: cf-app-utils 0.6 ruby lib

Gem::Specification.new do |s|
  s.name = "cf-app-utils"
  s.version = "0.6"

  s.required_rubygems_version = Gem::Requirement.new(">= 0") if s.respond_to? :required_rubygems_version=
  s.authors = ["Cloud Foundry"]
  s.date = "2016-05-25"
  s.description = "Helper methods for apps running on Cloud Foundry"
  s.email = ["vcap-dev@cloudfoundry.org"]
  s.files = ["lib/cf-app-utils", "lib/cf-app-utils/cf", "lib/cf-app-utils/cf/app", "lib/cf-app-utils/cf/app/credentials.rb", "lib/cf-app-utils/cf/app/environment.rb", "lib/cf-app-utils/cf/app/service.rb", "lib/cf-app-utils/cf.rb", "lib/cf-app-utils.rb", "LICENSE.txt", "README.md"]
  s.homepage = ""
  s.licenses = ["MIT"]
  s.require_paths = ["lib"]
  s.rubygems_version = "2.1.11"
  s.summary = "Helper methods for apps running on Cloud Foundry"

  if s.respond_to? :specification_version then
    s.specification_version = 4

    if Gem::Version.new(Gem::VERSION) >= Gem::Version.new('1.2.0') then
      s.add_development_dependency(%q<bundler>, ["~> 1.3"])
      s.add_development_dependency(%q<rake>, [">= 0"])
      s.add_development_dependency(%q<rspec>, [">= 0"])
    else
      s.add_dependency(%q<bundler>, ["~> 1.3"])
      s.add_dependency(%q<rake>, [">= 0"])
      s.add_dependency(%q<rspec>, [">= 0"])
    end
  else
    s.add_dependency(%q<bundler>, ["~> 1.3"])
    s.add_dependency(%q<rake>, [">= 0"])
    s.add_dependency(%q<rspec>, [">= 0"])
  end
end
