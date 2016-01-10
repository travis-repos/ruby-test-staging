require 'rspec/core/rake_task'

RSpec::Core::RakeTask.new

task default: [:spec, :oomded]

task :oomded do
  require 'securerandom'
  arr = []
  loop { arr << SecureRandom.uuid }
end
