/**
 *  @file
 *  @copyright defined in eos/LICENSE.txt
 */
#include "data_file_mng.hpp"
#include "data_file_mng.gen.hpp"

#include <eoslib/db.hpp>
#include <eoslib/types.hpp>
#include <eoslib/raw.hpp>


extern "C" {
    void init()  {
   }
   
    void apply( uint64_t code, uint64_t action ) {
        if( code == N(datafilemng) ) {
            if( action == N(datafilereg) ) {
                eosio::print("reg data file begin\n");
                auto req_Info = eosio::current_message<reg_data_file_req>();

                //eosio::require_auth( eosio::string_to_name(req_Info.info.user_name.get_data()) );               

                eosio::dump(req_Info);
                bytes b = eosio::raw::pack(req_Info.info);
                uint32_t err = store_str( N(datafilemng), N(datafileinfo), (char *)req_Info.file_hash.get_data(), req_Info.file_hash.get_size(), (char*)b.data, b.len);

                eosio::print("reg data file finished\n");           

            }     
            else {
                assert(0, "unknown message");
            }
        }
    }
}
