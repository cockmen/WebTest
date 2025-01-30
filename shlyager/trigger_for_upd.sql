CREATE TRIGGER update_rep_db_updated_at BEFORE UPDATE
    ON rep_db FOR EACH ROW EXECUTE PROCEDURE 
    update_updated_at_column();