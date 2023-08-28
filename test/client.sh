grpcurl  -d '{
    "Info": {
        "GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", 
        "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", 
        "BenefitDate": 1693195, 
        "TotalAmount": "500", 
        "UnsoldAmount": "100", 
        "TechniqueServiceFeeAmount": "50" 
    }
}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195,  "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500",  "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100"}}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement
grpcurl  -d '{"Info": {"ID": "b028dbec-fdcd-471c-b069-42ca3fb50476", "GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50" }}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement


grpcurl  -d '{"Infos": [{"GoodID": "87783d24-d188-4249-9378-a37c4e2c834a", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50"}]}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatements
grpcurl  -d '{"Infos": [{"GoodID": "eff4342e-3378-4f70-ae40-314d12c234a2", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50"},{"GoodID": "eff4342e-3378-4f70-ae40-314d12c234a2", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50"}]}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatements
grpcurl  -d '{"Infos": [{"GoodID": "eff4342e-3378-4f70-ae40-314d12c234a2", "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", "BenefitDate": 1693195, "TotalAmount": "500", "UnsoldAmount": "100", "TechniqueServiceFeeAmount": "50"}]}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatements

# delete
grpcurl  -d '{"Info": {"ID": "b028dbec-fdcd-471c-b069-42ca3fb50476"}}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.DeleteGoodStatement


grpcurl  -d '{
    "Info": {
        "GoodID": "020934d5-8b72-4707-9385-c683d9330633", 
        "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", 
        "BenefitDate": 1693195254, 
        "TotalAmount": "500", 
        "UnsoldAmount": "100", 
        "TechniqueServiceFeeAmount": "50" 
    }
}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement


grpcurl  -d '{
    "Info": {
        "GoodID": "8c3c6d7b-a7b9-48ad-ab3b-cf1fd4034cf3", 
        "CoinTypeID": "1f449b42-197a-4853-b6a7-04f1234fabf8", 
        "BenefitDate": 1693195554,
        "TotalAmount": "500",  
        "UnsoldAmount": "100", 
        "TechniqueServiceFeeAmount": "50" 
    }
}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.CreateGoodStatement

grpcurl  -d '{
    "Info": {
        "ID": "de6ccb28-520f-42b3-955f-15d76484a1a3"
    }
}' --plaintext 127.0.0.1:50421 ledger.middleware.good.ledger.statement.v2.Middleware.DeleteGoodStatement


grpcurl -d '{
    "Info":{
        "AppID":"ff2c5d50-be56-413e-aba5-9c7ad888a769",
        "UserID":"06094f12-0c0c-43d9-ae0a-f34064ce84ca",
        "CoinTypeID":"2b5f37c9-6a6a-4245-bdff-32281ae802e1",
        "IOType":"Incoming",
        "IOSubType":"Payment",
        "Amount":"1",
        "IOExtra":"{\"ff\":\"c8b73798-44c6-11ee-b797-222502ecd952\"}"
        }
}'  --plaintext localhost:50421 ledger.middleware.ledger.statement.v2.Middleware.CreateStatement

2、创建一条outcoming的payment statement
grpcurl --plaintext -d '{
    "Info":{
        "AppID":"ff2c5d50-be56-413e-aba5-9c7ad888a769",
        "UserID":"06094f12-0c0c-43d9-ae0a-f34064ce84ca",
        "CoinTypeID":"2b5f37c9-6a6a-4245-bdff-32281ae802e1",
        "IOType":"Outcoming",
        "IOSubType":"Payment",
        "Amount":"1",
        "IOExtra":"{\"ff\":\"c8b83ca6-44c6-11ee-983c-222502ecd952\"}"
    }
}' localhost:50421 ledger.middleware.ledger.statement.v2.Middleware.CreateStatement

3、删除该outcoming的statement记录
grpcurl --plaintext -d '{"Info":{"ID":"1ba2e541-6fe4-4fb6-94e4-ff28a0347967"}}' localhost:50421 ledger.middleware.ledger.statement.v2.Middleware.DeleteStatement