INSERT INTO `project_contract` (
  `id`,
  `name`,
  `description`,
  `img`,
  `project_website`,
  `status`,
  `amount`,
  `sale_contract_address`,
  `token_address`,
  `token_name`,
  `symbol`,
  `decimals`,
  `tge`,
  `current_price`,
  `registration_time_starts`,
  `registration_time_ends`,
  `number_of_registrants`,
  `sale_start`,
  `sale_end`,
  `token_price_in_PT`,
  `max_participation`,
  `total_tokens_sold`,
  `amount_of_tokens_to_sell`,
  `total_raised`,
  `earnings_withdrawn`,
  `leftover_withdrawn`,
  `unlock_time`,
  `create_time`,
  `update_time`
) VALUES (
  1,                                          -- id
  'Demo Project',                             -- name
  '这是一个测试项目',                          -- description
  'https://example.com/logo.png',             -- img
  'https://example.com',                      -- projectWebsite
  0,                                          -- status
  '1000000',                                  -- amount
  '0x1234567890abcdef1234567890abcdef12345678', -- saleContractAddress
  '0xabcdefabcdefabcdefabcdefabcdefabcdefabcd', -- tokenAddress
  'DemoToken',                                -- tokenName
  'DMT',                                      -- symbol
  18,                                         -- decimals
  '2026-03-15 12:00:00',                      -- tge
  100,                                        -- currentPrice
  '2026-03-10 00:00:00',                      -- registrationTimeStarts
  '2026-03-12 23:59:59',                      -- registrationTimeEnds
  1000,                                       -- number_of_registrants
  '2026-03-16 00:00:00',                      -- saleStart
  '2026-03-20 23:59:59',                      -- saleEnd
  '0.1',                                      -- tokenPriceInPT
  '1000',                                     -- maxParticipation
  '500000',                                   -- totalTokensSold
  '800000',                                   -- amountOfTokensToSell
  '50000',                                    -- totalRaised
  1,                                          -- earningsWithdrawn (true)
  0,                                          -- leftoverWithdrawn (false)
  '2026-03-25 00:00:00',                      -- unlockTime
  NOW(),                                      -- createTime
  NOW()                                       -- updateTime
);