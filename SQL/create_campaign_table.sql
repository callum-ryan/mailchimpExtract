CREATE TABLE campaign (
	id VARCHAR(255) NOT NULL PRIMARY KEY,
	list_id VARCHAR(255) REFERENCES list (id),
	segment_id VARCHAR(255) REFERENCES segment (id),
	winning_campaign_id VARCHAR(255) REFERENCES campaign (id),
	winning_combination_id VARCHAR(255) REFERENCES campaign_variate_combination (id),
	archive_url VARCHAR(255),
	authenticate BOOL,
	auto_footer BOOL,
	auto_tweet BOOL,
	clicktale VARCHAR(255),
	content_type VARCHAR(255),
	create_time VARCHAR(255),
	drag_and_drop BOOL,
	fb_comments BOOL,
	folder_id VARCHAR(255),
	from_name VARCHAR(255),
	google_analytics VARCHAR(255),
	inline_css BOOL,
	long_archive_url VARCHAR(255),
	reply_to VARCHAR(255),
	segment_text VARCHAR(255),
	send_time VARCHAR(255),
	status VARCHAR(255),
	subject_line VARCHAR(255),
	template_id INT,
	test_size INT,
	timewarp BOOL,
	title VARCHAR(255),
	to_name VARCHAR(255),
	track_ecomm_360 BOOL,
	track_goals BOOL,
	track_html_clicks BOOL,
	track_opens BOOL,
	track_text_clicks BOOL,
	type VARCHAR(255),
	use_conversation BOOL,
	wait_time INT,
	winner_criteria VARCHAR(255)
);