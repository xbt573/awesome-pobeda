gem_dir = File.expand_path('simpleidn', __dir__)

# Подменяем $LOAD_PATH на время загрузки
original_load_path = $LOAD_PATH.dup
begin
  $LOAD_PATH.replace([File.join(gem_dir, 'lib')])
  require 'simpleidn'
ensure
  $LOAD_PATH.replace(original_load_path)
end