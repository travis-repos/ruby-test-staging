require 'rspec/core/rake_task'

RSpec::Core::RakeTask.new

task default: [:spec, :oomded]

task :oomded do
  require 'securerandom'
  GC.disable
  arr = []
  s = SecureRandom.uuid
  loop do
    arr << "#{s}-#{arr.length}"
    print arr.length, ',' if arr.length % 10_000 == 0
  end
end
