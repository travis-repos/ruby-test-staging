require 'rspec/core/rake_task'

RSpec::Core::RakeTask.new

task default: [:spec, :oomded]

task :oomded do
  require 'securerandom'
  arr = []
  loop do
    arr << SecureRandom.uuid
    print '.' if arr.length % 9999 == 0
  end
end
