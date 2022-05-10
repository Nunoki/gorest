CREATE TRIGGER update_json_blob_mod_time BEFORE UPDATE ON json_blob FOR EACH ROW EXECUTE PROCEDURE  update_modified_column();
