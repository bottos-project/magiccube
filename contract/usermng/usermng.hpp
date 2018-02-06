/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */
#include <eoslib/eos.hpp>
#include <eoslib/string.hpp>

struct user_base_Info {
   uint8_t  user_type;  /* personal or company */
   eosio::string email;
   uint8_t   role_type; /* provider, consumer, arbiter, provider+consumer */
   eosio::string signature;
   eosio::string active_key;
   eosio::string owner_key;
};


/* @abi action insertkv2
 * @abi table
*/
struct add_user_req {
   eosio::string user_name;
   user_base_Info info;
};

