Raw Data loader for hotspot-store retrieving mass data from OpenApi


1.Argument 설정:
OpenAPI spec에 따라 data fetch 하는 paging 인자를 설정할 수 있으며,
configMap/hotstore-config 에 정의되어 있다.
인자 변경 시 pod/hotstore를 kill 하면 pod이 자동으로 재생성 되면서 새로운 인자를 load한다


2.Batch schedule설정:
cronJob/hotstore 에 정의된 cron스케쥴로 수행되며, 현재 매월 1일, 15일 3시 45분에 수행된다.
배치잡 수행 이력은 성공 1건/실패시 3건의 이력을 보관한다.
설정 변경 시 cronJob만 reload하면 된다.


3.RestAPI:
수동으로 데이터를 밀어 넣을 수 있도록 api endpoint 제공(TBD)


4.Slack integration:
배치 에러 발생 시 slack webhook을 통해 얼러팅을 보내준다
